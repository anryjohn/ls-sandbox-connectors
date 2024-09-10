package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	chpb "github.com/luthersystems/sandbox/api/chpb/v1"

	"github.com/sirupsen/logrus"
)

const (
	COOKIE_NAME = "SESSION_ID"
)

type CamundaInspectConnector struct {
	tripper     http.RoundTripper
	cookieURL   *url.URL
	logonFormat string
	queryFormat string
	tries       int

	lock    sync.Mutex
	token   string
	recency time.Time
}

type camundaInspectTripper struct {
}

func (s *camundaInspectTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == "POST" && req.URL.String() == "http://operate.byfn:8080/auth/login?username=demo&password=demo" {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header: http.Header(map[string][]string{
				"Content-Type": []string{"text/plain"},
				"Set-Cookie":   []string{fmt.Sprintf("%s=deadbeef", COOKIE_NAME)},
			}),
			Body:          ioutil.NopCloser(bytes.NewReader([]byte(""))),
			ContentLength: 0,
			Close:         false,
			Uncompressed:  true,
			Trailer:       map[string][]string{},
			Request:       req,
		}, nil
	}

	if req.Method == "GET" && strings.HasPrefix(req.URL.String(), "http://operate.byfn:8080/v1/process-instances/") {
		out := `{"state":"COMPLETED","other":"MISC"}`

		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header: http.Header(map[string][]string{
				"Content-Type": []string{"application/json"},
			}),
			Body:          ioutil.NopCloser(bytes.NewReader([]byte(out))),
			ContentLength: int64(len(out)),
			Close:         false,
			Uncompressed:  true,
			Trailer:       map[string][]string{},
			Request:       req,
		}, nil
	}

	return &http.Response{
		Status:     "404 Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header(map[string][]string{
			"Content-Type": []string{"text/plain"},
		}),
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(""))),
		ContentLength: 0,
		Close:         false,
		Uncompressed:  true,
		Trailer:       map[string][]string{},
		Request:       req,
	}, nil
}

func NewCamundaInspectConnector() (*CamundaInspectConnector, error) {
	var err error
	s := &CamundaInspectConnector{
		tripper: &camundaInspectTripper{}, // TODO: remove this tripper mock when Operate can actually authenticate
	}
	s.cookieURL, err = url.Parse("http://operate.byfn:8080/")
	if err != nil {
		return nil, err
	}
	s.logonFormat = "http://operate.byfn:8080/auth/login?username=demo&password=demo"
	s.queryFormat = "http://operate.byfn:8080/v1/process-instances/%d"
	s.tries = 60 / 5
	s.token = "invalid"
	s.recency = time.Now().Add(-time.Minute)
	return s, nil
}

func (s *CamundaInspectConnector) inspectToken() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.token
}

func (s *CamundaInspectConnector) refreshToken() {
	if (func() bool {
		s.lock.Lock()
		defer s.lock.Unlock()

		return time.Since(s.recency) >= time.Minute
	})() {
		hereby := time.Now()

		jar, err := cookiejar.New(&cookiejar.Options{})
		if err != nil {
			logrus.Info("failed to create cookie jar")
			return
		}

		client := &http.Client{
			Transport: s.tripper,
			Jar:       jar,
		}

		token := "invalid"

		resp, err := client.Get(s.logonFormat)
		if err != nil {
			logrus.Info("failed to authenticate, trying again soon")
			return
		}

		resp.Body.Close()

		url, err := url.Parse(fmt.Sprintf(s.queryFormat, 0))
		if err != nil {
			logrus.Info("failed to create URL")
			return
		}

		for _, cookie := range jar.Cookies(url) {
			if cookie.Name == COOKIE_NAME {
				token = cookie.Value
			}
		}

		s.lock.Lock()
		defer s.lock.Unlock()

		if hereby.UnixNano() > s.recency.UnixNano() {
			s.token = token
			s.recency = hereby
		}
	}
}

type InterestingJSON struct {
	State string `json:"state"`
}

func (s *CamundaInspectConnector) Handle(ctx context.Context, req *chpb.CamundaInspectRequest) (*chpb.CamundaInspectResponse, error) {
	// define a helper function that sleeps only when the client
	// is attempting to perform long polling
	//
	// this is used to inject delay on paths that would typically
	// retry in order to avoid flooding
	slowly := (func() {
		if req.WaitForState != "" {
			time.Sleep(5 * time.Second)
		}
	})

	fmt.Fprintf(os.Stderr, "HANDLE: %d\n", req.GetProcessInstanceKey())

	var body []byte

	refund := true

	for i := 0; i < s.tries; i++ {
		token := s.inspectToken()

		jar, err := cookiejar.New(&cookiejar.Options{})
		if err != nil {
			slowly()

			return &chpb.CamundaInspectResponse{
				Success:    false,
				Diagnostic: err.Error(),
			}, nil
		}

		client := &http.Client{
			Transport: s.tripper,
			Jar:       jar,
		}

		client.Jar.SetCookies(s.cookieURL, []*http.Cookie{
			&http.Cookie{
				Name:  COOKIE_NAME,
				Value: token,
			},
		})

		resp, err := client.Get(fmt.Sprintf(s.queryFormat, req.GetProcessInstanceKey()))
		if err != nil {
			slowly()

			return &chpb.CamundaInspectResponse{
				Success:    false,
				Diagnostic: err.Error(),
			}, nil
		}

		if resp.StatusCode == 401 {
			s.refreshToken()

			if refund {
				i -= 1
				refund = false
			} else {
				time.Sleep(5 * time.Second)
				continue
			}
		}

		if resp.StatusCode != 200 {
			slowly()

			return &chpb.CamundaInspectResponse{
				Success:    false,
				Diagnostic: fmt.Sprintf("CONNR: non-200 status: %s", resp.Status),
			}, nil
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			slowly()

			return &chpb.CamundaInspectResponse{
				Success:    false,
				Diagnostic: err.Error(),
			}, nil
		}

		var interesting InterestingJSON
		err = json.Unmarshal(body, &interesting)
		if err != nil {
			slowly()

			return &chpb.CamundaInspectResponse{
				Success:    false,
				Diagnostic: err.Error(),
			}, nil
		}

		if req.WaitForState == "" || (interesting.State == req.WaitForState) {
			return &chpb.CamundaInspectResponse{
				Success: true,
				Content: hex.EncodeToString(body),
			}, nil
		}

		time.Sleep(5 * time.Second)
	}

	return &chpb.CamundaInspectResponse{
		Success: true,
		Content: hex.EncodeToString(body),
	}, nil
}

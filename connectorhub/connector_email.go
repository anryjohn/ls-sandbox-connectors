package main

import (
	"context"
	"net/smtp"
	"strings"

	chpb "github.com/luthersystems/sandbox/api/chpb/v1"
)

type EmailConnector struct {
	sendmail func(recipient string, title string, body string) error
}

func NewEmailConnector() (*EmailConnector, error) {
	s := &EmailConnector{}
	sender := "marketing@example.com"
	s.sendmail = (func(recipient string, title string, body string) error {
		body = strings.NewReplacer("\n", "\r\n").Replace(body)
		body = "To: " + recipient + "\r\nFrom: " + sender + "\r\nSubject: " + title + "\r\n\r\n" + body + "\r\n\r\n"
		return smtp.SendMail("demo-mta-1.byfn:25", nil, sender, []string{recipient}, []byte(body))
	})
	return s, nil
}

func (s *EmailConnector) Handle(ctx context.Context, req *chpb.EmailRequest) (*chpb.EmailResponse, error) {
	err := s.sendmail(req.GetRecipient(), req.GetTitle(), req.GetBody())
	if err != nil {
		return nil, err
	}
	return &chpb.EmailResponse{}, nil
}

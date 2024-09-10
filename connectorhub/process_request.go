package main

import (
	"context"
	"encoding/json"
	"fmt"

	chpb "github.com/luthersystems/sandbox/api/chpb/v1"

	"github.com/luthersystems/sandbox/connectorhub/internal/events"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

type connectors struct {
	postgres       *PostgresConnector
	email          *EmailConnector
	camundaStart   *CamundaStartConnector
	camundaInspect *CamundaInspectConnector
}

var globalConnectors *connectors

func processRequestInitialize() (*connectors, error) {
	postgres, err := NewPostgresConnector()
	if err != nil {
		panic(err)
	}

	email, err := NewEmailConnector()
	if err != nil {
		panic(err)
	}

	camundaStart, err := NewCamundaStartConnector()
	if err != nil {
		panic(err)
	}

	camundaInspect, err := NewCamundaInspectConnector()
	if err != nil {
		panic(err)
	}

	return &connectors{
		postgres:       postgres,
		email:          email,
		camundaStart:   camundaStart,
		camundaInspect: camundaInspect,
	}, nil
}

func processRequest(ctx context.Context, req json.RawMessage, reqErr error) (json.RawMessage, error) {
	if globalConnectors == nil {
		var err error
		globalConnectors, err = processRequestInitialize()
		if err != nil {
			panic(err)
		}
	}

	if reqErr != nil {
		return nil, fmt.Errorf("request had error: %w", reqErr)
	}

	event := &chpb.Event{}
	if err := protojson.Unmarshal([]byte(req), event); err != nil {
		return nil, err
	}
	if postgres := event.GetConnectorPostgres(); postgres != nil {
		resp, err := globalConnectors.postgres.Handle(ctx, postgres)
		if err != nil {
			return nil, err
		}
		bytes, err := (&protojson.MarshalOptions{UseProtoNames: true}).Marshal(resp)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(bytes), nil
	} else if email := event.GetConnectorEmail(); email != nil {
		resp, err := globalConnectors.email.Handle(ctx, email)
		if err != nil {
			return nil, err
		}
		bytes, err := (&protojson.MarshalOptions{UseProtoNames: true}).Marshal(resp)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(bytes), nil
	} else if camundaStart := event.GetConnectorCamundaStart(); camundaStart != nil {
		resp, err := globalConnectors.camundaStart.Handle(ctx, camundaStart)
		if err != nil {
			return nil, err
		}
		bytes, err := (&protojson.MarshalOptions{UseProtoNames: true}).Marshal(resp)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(bytes), nil
	} else if camundaInspect := event.GetConnectorCamundaInspect(); camundaInspect != nil {
		resp, err := globalConnectors.camundaInspect.Handle(ctx, camundaInspect)
		if err != nil {
			return nil, err
		}
		bytes, err := (&protojson.MarshalOptions{UseProtoNames: true}).Marshal(resp)
		if err != nil {
			return nil, err
		}
		return json.RawMessage(bytes), nil
	} else {
		// sometimes we get a bare request_id with no connector indicated; I believe this is due to the claims example
		return nil, fmt.Errorf("unknown connector (req=%v)", string(req))
	}
}

func (s *g) Run() error {
	if s.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	ctx := s.ctx

	var gatewayOpts []events.Option
	if s.StartBlockNumber > 0 {
		gatewayOpts = append(gatewayOpts, events.WithStartBlock(uint64(s.StartBlockNumber)))
	}
	if s.CheckpointFile != "" {
		gatewayOpts = append(gatewayOpts, events.WithCheckpointFile(s.CheckpointFile))
	}
	logrus.WithContext(ctx).Info("connecting to gateway")

	stream, err := events.GatewayEvents(gatewayCfg, gatewayOpts...)
	if err != nil {
		return fmt.Errorf("gateway events: %w", err)
	}

	ctx, cancel := context.WithCancel(s.ctx)

	go func() {
		logrus.WithContext(ctx).Info("listening for events")
		for {
			select {
			case event := <-stream.Listen():
				if event == nil {
					logrus.WithContext(ctx).Info("nil event (stale checkpoint file?), exiting...")
					return
				}
				req, err := event.RequestBody()
				if err != nil {
					logrus.WithContext(ctx).WithError(err).Error("event received with error")
				}
				go (func() {
					resp, err := processRequest(ctx, req, err)
					if err := event.Callback(resp, err); err != nil {
						logrus.WithContext(ctx).WithError(err).Error("event callback failed")
					} else {
						logrus.WithContext(ctx).Info("callback successful")
					}
				})()
			case <-ctx.Done():
				logrus.WithContext(ctx).Info("event listener shutting down...")
				return
			}
		}
	}()

	<-ctx.Done()
	logrus.WithContext(ctx).Info("signal handler called")
	cancel()
	if err := stream.Done(); err != nil {
		logrus.WithError(err).Debug("stream done")
	}

	logrus.WithContext(ctx).Info("connectorhub exited!")

	return nil
}

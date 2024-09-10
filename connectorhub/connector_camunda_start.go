package main

import (
	"context"
	"encoding/hex"
	"encoding/json"

	chpb "github.com/luthersystems/sandbox/api/chpb/v1"

	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

type CamundaStartConnector struct {
	client zbc.Client
}

func NewCamundaStartConnector() (*CamundaStartConnector, error) {
	s := &CamundaStartConnector{}
	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress: "zeebe.byfn:26500",
		UsePlaintextConnection: true,
	})
	if err != nil {
		return nil, err
	}
	s.client = client
	return s, nil
}

func (s *CamundaStartConnector) Handle(ctx context.Context, req *chpb.CamundaStartRequest) (*chpb.CamundaStartResponse, error) {
	var variables map[string]interface{}

	blob, err := hex.DecodeString(req.GetVariables())
	if err != nil {
		return &chpb.CamundaStartResponse{
			Success: false,
			Diagnostic: err.Error(),
		}, nil
	}

	err = json.Unmarshal(blob, &variables)
	if err != nil {
		return &chpb.CamundaStartResponse{
			Success: false,
			Diagnostic: err.Error(),
		}, nil
	}

	command, err := s.client.NewCreateInstanceCommand().BPMNProcessId(req.GetProcessId()).LatestVersion().VariablesFromMap(variables)
	if err != nil {
		return &chpb.CamundaStartResponse{
			Success: false,
			Diagnostic: err.Error(),
		}, nil
	}

	process, err := command.Send(ctx)
	if err != nil {
		return &chpb.CamundaStartResponse{
			Success: false,
			Diagnostic: err.Error(),
		}, nil
	}

	return &chpb.CamundaStartResponse{
		Success: true,
		ProcessInstanceKey: process.GetProcessInstanceKey(),
	}, nil
}

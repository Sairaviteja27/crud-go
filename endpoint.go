package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// For each method, we define request and response structs
type messageRequest struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type createMessageResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCreateMessageEndpoint(svc RecordService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(messageRequest)
		v, err := svc.addRecord(ctx, req)
		if err != nil {
			return createMessageResponse{v, err.Error()}, nil
		}
		return createMessageResponse{v, ""}, nil
	}
}

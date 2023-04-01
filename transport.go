package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty message")

func decodeAddMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request messageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Message == "" {
		return nil, ErrEmpty
	}
	return request, nil
}

func decodeReadMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request messageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Message == "" {
		return nil, ErrEmpty
	}
	return request, nil
}
func decodeUpdateMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request messageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Message == "" {
		return nil, ErrEmpty
	}
	return request, nil
}
func decodeDeleteMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request messageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Message == "" {
		return nil, ErrEmpty
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

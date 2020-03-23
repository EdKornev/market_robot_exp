package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type getStatusRequest struct{}

type getStatusResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type connectRequest struct{}

type connectResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type disconnectRequest struct{}

type disconnectResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

func decodeGetStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getStatusRequest
	return req, nil
}

func decodeConnectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req connectRequest
	return req, nil
}

func decodeDisconnectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req disconnectRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

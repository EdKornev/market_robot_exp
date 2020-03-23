package api

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetStatusEndpoint  endpoint.Endpoint
	ConnectEndpoint    endpoint.Endpoint
	DisconnectEndpoint endpoint.Endpoint
}

func MakeGetStatusEndpoint(srv Api) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getStatusRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetStatus(ctx)
		if err != nil {
			return getStatusResponse{d != -1, err.Error()}, nil
		}
		return getStatusResponse{d != -1, ""}, nil
	}
}

func MakeConnectEndpoint(srv Api) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(connectRequest) // we really just need the request, we don't use any value from it
		b, err := srv.Connect(ctx)
		if err != nil {
			return connectResponse{b, err.Error()}, err
		}
		return connectResponse{b, ""}, nil
	}
}

func MakeDisconnectEndpoint(srv Api) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(disconnectRequest)
		b, err := srv.Disconnect(ctx)
		if err != nil {
			return disconnectResponse{b, err.Error()}, nil
		}
		return disconnectResponse{b, ""}, nil
	}
}

func (e Endpoints) GetStatus(ctx context.Context) (bool, error) {
	req := getStatusRequest{}
	resp, err := e.GetStatusEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	getResp := resp.(getStatusResponse)
	if getResp.Err != "" {
		return false, errors.New(getResp.Err)
	}
	return getResp.Valid, nil
}

func (e Endpoints) Disconnect(ctx context.Context) (bool, error) {
	req := disconnectRequest{}
	resp, err := e.DisconnectEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	statusResp := resp.(disconnectResponse)
	return statusResp.Valid, nil
}

func (e Endpoints) Connect(ctx context.Context, date string) (bool, error) {
	req := connectRequest{}
	resp, err := e.ConnectEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(connectResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}

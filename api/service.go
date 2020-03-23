package api

import "context"

type Api interface {
	Connect(ctx context.Context) (bool, error)
	Disconnect(ctx context.Context) (bool, error)
	GetStatus(ctx context.Context) (int, error)
}

type dataApi struct{}

func ApiImpl() Api {
	return dataApi{}
}

func (dataApi) Connect(ctx context.Context) (bool, error) {
	return true, nil
}

func (dataApi) Disconnect(ctx context.Context) (bool, error) {
	return true, nil
}

func (dataApi) GetStatus(ctx context.Context) (int, error) {
	return 0, nil
}

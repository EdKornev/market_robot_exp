package api

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	r.Methods("GET").Path("/get_status").Handler(httptransport.NewServer(
		endpoints.GetStatusEndpoint,
		decodeGetStatusRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/connect").Handler(httptransport.NewServer(
		endpoints.ConnectEndpoint,
		decodeConnectRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/disconnect").Handler(httptransport.NewServer(
		endpoints.DisconnectEndpoint,
		decodeDisconnectRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

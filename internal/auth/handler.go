package auth

import (
	"net/http"

	connectrpc "buf.build/gen/go/salih-g/gulbeyaz/connectrpc/go/auth/v1/authv1connect"
	authv1 "buf.build/gen/go/salih-g/gulbeyaz/protocolbuffers/go/auth/v1"
	"connectrpc.com/connect"
	"golang.org/x/net/context"
)

type Handler struct {
	connectrpc.UnimplementedAuthServiceHandler
}

func (h *Handler) RegisterRoutes() (string, http.Handler) {
	return connectrpc.NewAuthServiceHandler(
		h,
	)
}

func (h *Handler) Login(
	ctx context.Context,
	request *connect.Request[authv1.LoginRequest],
) (*connect.Response[authv1.LoginResponse], error) {

	return nil, nil
}

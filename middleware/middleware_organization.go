package middleware

import (
	"context"
	"errors"

	"connectrpc.com/connect"
)

func NewOrganizationInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if !req.Spec().IsClient {
				if req.Header().Get("Authorization") == "" {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no token provided"))
				}
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

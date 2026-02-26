package middleware

import (
	"app/constant"
	"app/pkg"
	"context"
	"strings"

	"connectrpc.com/connect"
)

type principalContextKey struct{}

type Principal struct {
	Token pkg.MetaToken
}

func PrincipalFromContext(ctx context.Context) (Principal, bool) {
	principal, ok := ctx.Value(principalContextKey{}).(Principal)
	return principal, ok
}

func NewAuthInterceptor(publicProcedures map[string]struct{}) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			if _, isPublic := publicProcedures[req.Spec().Procedure]; isPublic {
				return next(ctx, req)
			}

			header := req.Header().Get("Authorization")
			token, ok := parseBearerToken(header)
			if !ok {
				return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			}

			metaToken, err := pkg.VerifyTokenHeader(token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			}
			if metaToken.TokenType != "access" {
				return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			}

			ctx = context.WithValue(ctx, principalContextKey{}, Principal{Token: metaToken})
			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

func parseBearerToken(authorizationHeader string) (string, bool) {
	parts := strings.Fields(authorizationHeader)
	if len(parts) != 2 {
		return "", false
	}
	if !strings.EqualFold(parts[0], "Bearer") {
		return "", false
	}
	if parts[1] == "" {
		return "", false
	}
	return parts[1], true
}

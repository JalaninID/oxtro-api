package middleware

import (
	toolsv1 "app/gen/tools/v1"
	"app/pkg"
	"context"
	"testing"

	"connectrpc.com/connect"
)

func TestAuthInterceptor_RejectsMissingAuthorization(t *testing.T) {
	interceptor := NewAuthInterceptor(nil)
	req := connect.NewRequest(&toolsv1.Empty{})

	next := interceptor(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		return connect.NewResponse(&toolsv1.Empty{}), nil
	})

	_, err := next(context.Background(), req)
	if err == nil {
		t.Fatalf("expected unauthenticated error when authorization is missing")
	}
}

func TestAuthInterceptor_AllowsValidBearerToken(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-secret-key")
	t.Setenv("JWT_TIME_DURATION", "15")
	token, err := pkg.Sign(map[string]any{
		"id":         "user-123",
		"session_id": "5",
		"token_type": "access",
	}, 5)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	interceptor := NewAuthInterceptor(nil)
	req := connect.NewRequest(&toolsv1.Empty{})
	req.Header().Set("Authorization", "Bearer "+token)

	next := interceptor(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		principal, ok := PrincipalFromContext(ctx)
		if !ok {
			t.Fatalf("expected principal in context")
		}
		if principal.Token.ID != "user-123" {
			t.Fatalf("expected principal user-123, got %s", principal.Token.ID)
		}
		return connect.NewResponse(&toolsv1.Empty{}), nil
	})

	if _, err := next(context.Background(), req); err != nil {
		t.Fatalf("expected request to pass auth interceptor, got %v", err)
	}
}

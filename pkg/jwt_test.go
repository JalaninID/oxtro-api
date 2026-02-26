package pkg

import (
	"os"
	"testing"
	"time"
)

func TestSignAndVerifyTokenHeader(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-secret-key")
	t.Setenv("JWT_TIME_DURATION", "15")

	token, err := Sign(map[string]any{
		"id":         "user-uuid",
		"session_id": "1",
		"token_type": "access",
	}, 1)
	if err != nil {
		t.Fatalf("unexpected sign error: %v", err)
	}

	meta, err := VerifyTokenHeader(token)
	if err != nil {
		t.Fatalf("unexpected verify error: %v", err)
	}
	if meta.ID != "user-uuid" {
		t.Fatalf("expected id user-uuid, got %s", meta.ID)
	}
	if meta.SessionID != "1" {
		t.Fatalf("expected session id 1, got %s", meta.SessionID)
	}
	if meta.TokenType != "access" {
		t.Fatalf("expected token_type access, got %s", meta.TokenType)
	}
	if meta.Exp <= time.Now().Unix() {
		t.Fatalf("expected token expiration in future, got %d", meta.Exp)
	}
}

func TestVerifyTokenHeader_InvalidSecret(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-secret-key")
	t.Setenv("JWT_TIME_DURATION", "15")

	token, err := Sign(map[string]any{"id": "user-uuid", "token_type": "access"}, 1)
	if err != nil {
		t.Fatalf("unexpected sign error: %v", err)
	}

	if err := os.Setenv("JWT_SECRET_KEY", "another-secret"); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	if _, err := VerifyTokenHeader(token); err == nil {
		t.Fatalf("expected verify to fail with wrong secret")
	}
}

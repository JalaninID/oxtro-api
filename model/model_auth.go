package model

import "time"

type AuthSession struct {
	ID            int
	UUID          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	UserID        int
	TokenHash     string
	DeviceID      string
	UserAgent     string
	IPAddress     string
	ExpiresAt     time.Time
	RevokedAt     *time.Time
	RotatedFromID *int
}

func (s *AuthSession) TableName() string {
	return "auth_sessions"
}

type PasswordReset struct {
	ID        int
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int
	TokenHash string
	ExpiresAt time.Time
	UsedAt    *time.Time
}

func (s *PasswordReset) TableName() string {
	return "password_resets"
}

type EmailVerification struct {
	ID        int
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int
	TokenHash string
	ExpiresAt time.Time
	UsedAt    *time.Time
}

func (s *EmailVerification) TableName() string {
	return "email_verifications"
}

type LoginAttempt struct {
	ID        int
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int
	Email     string
	IPAddress string
	Success   bool
	Reason    string
}

func (s *LoginAttempt) TableName() string {
	return "login_attempts"
}

type SecurityEvent struct {
	ID        int
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int
	EventType string
	Severity  string
	Metadata  string
}

func (s *SecurityEvent) TableName() string {
	return "security_events"
}

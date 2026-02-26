package pkg

import "testing"

func TestHashAndComparePassword(t *testing.T) {
	t.Setenv("BCRYPT_COST", "4")

	password := "StrongPass123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected hash error: %v", err)
	}
	if hash == password {
		t.Fatalf("expected hash to differ from password")
	}

	if err := ComparePassword(hash, password); err != nil {
		t.Fatalf("expected password comparison to succeed, got %v", err)
	}
}

func TestValidatePassword(t *testing.T) {
	if err := ValidatePassword("weak"); err == nil {
		t.Fatalf("expected weak password to fail validation")
	}
	if err := ValidatePassword("StrongPass123"); err != nil {
		t.Fatalf("expected strong password to pass validation, got %v", err)
	}
}

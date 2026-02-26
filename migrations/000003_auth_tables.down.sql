DROP TABLE IF EXISTS security_events;
DROP TABLE IF EXISTS login_attempts;
DROP TABLE IF EXISTS email_verifications;
DROP TABLE IF EXISTS password_resets;
DROP TABLE IF EXISTS auth_sessions;

ALTER TABLE users
    DROP COLUMN IF EXISTS email_verified_at,
    DROP COLUMN IF EXISTS locked_until,
    DROP COLUMN IF EXISTS is_active;

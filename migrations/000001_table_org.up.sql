CREATE TABLE IF NOT EXISTS "organizations"(
    "id" BIGSERIAL PRIMARY KEY, -- Auto Increment
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "uuid" VARCHAR(36) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "domain" VARCHAR(255) NOT NULL,
    "logo" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_organization"(
    "id" BIGSERIAL PRIMARY KEY, -- Auto Increment
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "user_id" BIGINT NOT NULL,
    "organization_id" BIGINT NOT NULL,
    "role_id" BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS "users"(
    "id" BIGSERIAL PRIMARY KEY, -- Auto Increment
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "uuid" VARCHAR(36) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "avatar" VARCHAR(255) NULL,
    "bio" TEXT NULL,
    "password" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "roles"(
    "id" BIGSERIAL PRIMARY KEY, -- Auto Increment
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "uuid" VARCHAR(36) NOT NULL,
    "organization_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "access" JSONB NULL
);
CREATE TABLE IF NOT EXISTS "organization_configs"(
    "id" BIGSERIAL PRIMARY KEY, -- Auto Increment
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "uuid" VARCHAR(36) NOT NULL,
    "organization_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "enable" BOOLEAN NOT NULL DEFAULT FALSE,
    "description" TEXT NOT NULL
);
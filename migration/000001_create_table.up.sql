CREATE TABLE "users"(
    "id" BIGSERIAL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "username" VARCHAR(100) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "profile" TEXT NOT NULL,
    "status" VARCHAR(25) NOT NULL,
    "is_premium" BOOLEAN NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");

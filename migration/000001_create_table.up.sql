CREATE TABLE "users"(
    "id" VARCHAR(36) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "username" VARCHAR(100) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) CHECK
      ("gender" IN('FEMALE', 'MALE')) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "profile" TEXT NOT NULL,
    "status" VARCHAR(25) NOT NULL,
    "is_premium" BOOLEAN NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "interests"(
    "id" VARCHAR(36) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "user_id" VARCHAR(36) NOT NULL,
    "is_interest" BOOLEAN NOT NULL,
    "interest_user_id" VARCHAR(36) NOT NULL
);
ALTER TABLE
    "interests" ADD PRIMARY KEY("id");
-- create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "email" character varying(255) NULL, "password" character varying(255) NULL, "last_login_at" timestamp NULL, "last_login_ip" character varying(255) NULL, "permission" bigint NULL DEFAULT 0, "is_banned" boolean NULL DEFAULT false, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "delete_at" timestamptz NULL, PRIMARY KEY ("id"));
-- create index "idx_users_delete_at" to table: "users"
CREATE INDEX "idx_users_delete_at" ON "public"."users" ("delete_at");
-- create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");

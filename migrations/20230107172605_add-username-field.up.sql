-- modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "username" character varying(255) NULL;
-- create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "public"."users" ("username");

-- reverse: create index "idx_users_email" to table: "users"
DROP INDEX "public"."idx_users_email";
-- reverse: create index "idx_users_delete_at" to table: "users"
DROP INDEX "public"."idx_users_delete_at";
-- reverse: create "users" table
DROP TABLE "public"."users";

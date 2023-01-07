-- reverse: create index "idx_users_username" to table: "users"
DROP INDEX "public"."idx_users_username";
-- reverse: modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "username";

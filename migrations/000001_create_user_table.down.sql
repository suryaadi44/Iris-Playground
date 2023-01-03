BEGIN;

DROP INDEX IF EXISTS idx_users_delete_at;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;

DROP TABLE IF EXISTS public.users;

COMMIT;
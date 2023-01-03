BEGIN;

CREATE TABLE IF NOT EXISTS public.users (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email CHARACTER VARYING(255),
  password CHARACTER VARYING(255),
  last_login_at TIMESTAMP WITH TIME ZONE,
  last_login_ip CHARACTER VARYING(255),
  permission INT DEFAULT 0,
  is_banned BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
  delete_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_users_delete_at ON USERS USING BTREE (delete_at);
CREATE UNIQUE INDEX idx_users_email ON USERS USING BTREE (email);

COMMIT;
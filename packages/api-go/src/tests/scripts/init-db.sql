CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id uuid PRIMARY KEY,
	email varchar(255) UNIQUE NOT NULL,
	password_hash varchar NOT NULL,
	name varchar(255) NOT NULL,
	bio text,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.posts (
	id uuid PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
	content text NOT NULL,
	image_url varchar(255),
	is_public bool DEFAULT true,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.post_likes (
	id uuid PRIMARY KEY,
	post_id uuid NOT NULL REFERENCES public.posts(id) ON DELETE CASCADE,
	user_id uuid NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT unique_like_per_user UNIQUE (post_id, user_id)
);

-- Seed a test user
INSERT INTO users (id, email, password_hash, name)
VALUES ('3d0aef0a-0167-4dde-b013-75889f0ce8a3', 'test@example.com', '$argon2id$v=19$m=19456,t=2,p=1$QTdkZFRuTzYyc2UxZTJBMw$kc6kTP+WCxi9PYSSvmPmv3wSizG62ueIDidCjZtLuss', 'Tester');
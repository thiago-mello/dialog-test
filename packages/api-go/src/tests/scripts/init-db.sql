CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id uuid NOT NULL,
	email varchar(255) NOT NULL,
	password_hash varchar NOT NULL,
	"name" varchar(255) NOT NULL,
	bio text NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


CREATE TABLE public.posts (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	"content" text NOT NULL,
	image_url varchar(255) NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	is_public bool DEFAULT true NOT NULL,
	CONSTRAINT posts_pkey PRIMARY KEY (id),
	CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX posts_id_user_id_idx ON public.posts USING btree (id, user_id);


CREATE TABLE public.post_likes (
	id uuid NOT NULL,
	post_id uuid NOT NULL,
	user_id uuid NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT post_likes_pkey PRIMARY KEY (id),
	CONSTRAINT unique_like_per_user UNIQUE (post_id, user_id),
	CONSTRAINT post_likes_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT post_likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX post_likes_post_id_idx ON public.post_likes USING btree (post_id);

-- Seed a test user
INSERT INTO users (id, email, password_hash, name)
VALUES ('3d0aef0a-0167-4dde-b013-75889f0ce8a3', 'test@example.com', '$argon2id$v=19$m=19456,t=2,p=1$QTdkZFRuTzYyc2UxZTJBMw$kc6kTP+WCxi9PYSSvmPmv3wSizG62ueIDidCjZtLuss', 'Tester');
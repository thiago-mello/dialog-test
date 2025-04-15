{{define "post.Save"}}
    INSERT INTO posts (id, user_id, content, is_public, created_at, updated_at)
    VALUES (:id, :user_id, :content, :is_public, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    ON CONFLICT (id) DO UPDATE SET
        content = EXCLUDED.content,
        is_public = EXCLUDED.is_public,
        updated_at = CURRENT_TIMESTAMP
{{end}}
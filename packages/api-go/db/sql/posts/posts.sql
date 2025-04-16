{{define "post.Save"}}
    INSERT INTO posts (id, user_id, content, is_public, created_at, updated_at)
    VALUES (:id, :user_id, :content, :is_public, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    ON CONFLICT (id) DO UPDATE SET
        content = EXCLUDED.content,
        is_public = EXCLUDED.is_public,
        updated_at = CURRENT_TIMESTAMP
{{end}}

{{define "post.Update"}}
    UPDATE posts 
    SET 
        content = :content,
        is_public = :is_public,
        updated_at = CURRENT_TIMESTAMP
    WHERE id = :id AND user_id = :user_id
{{end}}

{{define "post.FindByID"}}
    SELECT
        id,
        user_id,
        content,
        is_public,
        created_at,
        updated_at
    FROM posts
    WHERE id = :id
{{end}}
{{define "post.Like"}}
    INSERT INTO post_likes (id, post_id, user_id, created_at)
    VALUES (:id, :post_id, :user_id, CURRENT_TIMESTAMP)
    ON CONFLICT (post_id, user_id) DO NOTHING
{{end}}

{{define "post.Unlike"}}
    DELETE FROM post_likes
    WHERE post_id = :post_id AND user_id = :user_id
{{end}}
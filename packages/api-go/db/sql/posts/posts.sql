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

{{define "post.ListPosts"}}
SELECT
    p.id,
    p."content",
    p.user_id,
    u."name" AS user_name,
    u.bio AS user_bio,
    p.created_at,
    p.updated_at,
    (
    SELECT
        count(*)
    FROM
        post_likes pl
    WHERE
        pl.post_id = p.id) AS likes,
    EXISTS (
    SELECT
        id
    FROM
        post_likes pl
    WHERE
        pl.post_id = p.id
        AND pl.user_id = :current_user_id) AS user_liked
FROM
    posts p
INNER JOIN users u ON
    p.user_id = u.id
WHERE
    p.is_public IS TRUE
    {{- if .LastSeenId}}
        AND p.id < :last_seen_id
    {{end}}
    {{- if .UserId}}
        AND p.user_id = :user_id
    {{end}}
ORDER BY
    p.id DESC
    LIMIT :page_size
{{end}}
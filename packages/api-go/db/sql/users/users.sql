{{define "user.Insert"}}
INSERT INTO users (id, email, password_hash, "name", bio, created_at, updated_at)
VALUES(:id, :email, :password_hash, :name, :bio, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
{{end}}

{{define "user.ExistsByEmail"}}
SELECT
    u.email
FROM
    users u
WHERE
    u.email = :email
{{end}}
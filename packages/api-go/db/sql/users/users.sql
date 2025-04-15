{{define "user.Insert"}}
    INSERT INTO users (id, email, password_hash, "name", bio, created_at, updated_at)
    VALUES(:id, :email, :password_hash, :name, :bio, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
{{end}}

{{define "user.FindByEmail"}}
    SELECT
        u.id,
        u.email,
        u.password_hash,
        u."name",
        u.bio,
        u.created_at,
        u.updated_at
    FROM
        users u
    WHERE
        u.email = :email
{{end}}

{{define "user.UpdateById"}}
    UPDATE
        users u
    SET
        "name" = :name,
        email = :email,
        bio = :bio,
        {{- if .PasswordHash}}
            password_hash = :password_hash,
        {{end}}
        updated_at = current_timestamp
    WHERE
        u.id = :id
{{end}}
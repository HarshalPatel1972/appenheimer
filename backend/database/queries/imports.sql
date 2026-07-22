-- name: CreateImport :one
INSERT INTO imports (
    id, source, status, started_at
) VALUES (
    $1, $2, $3, NOW()
) RETURNING *;

-- name: UpdateImport :one
UPDATE imports SET
    status = $2,
    finished_at = NOW(),
    report = $3
WHERE id = $1
RETURNING *;

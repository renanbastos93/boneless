-- name: ListExamples :many
SELECT * FROM examples;

-- name: GetExampleById :one
SELECT * FROM examples
WHERE id = ?;

-- name: CreateExample :execresult
INSERT INTO examples (
  id, created_at, message
) VALUES (
  ?, ?, ?
);

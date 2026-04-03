
-- name: GetBookByID :one
SELECT id, title, author, num_pages, rating from books
WHERE id = $1;

-- name: InsertBook :one
INSERT INTO books(title, author, num_pages, rating)
VALUES($1,$2,$3,$4)
RETURNING *;

-- name: ListAllBooks :many
SELECT id, title, author, num_pages, rating from books
ORDER BY id;

-- name: ListBooks :many
SELECT id, title, author, num_pages, rating from books
WHERE id>= $1 AND id<=$2
ORDER BY id;

-- name: DeleteBook :execrows
DELETE FROM books
Where id = $1;

-- name: FullUpdateBook :execrows
UPDATE books
SET title = $2, author = $3, num_pages = $4, rating = $5
WHERE id = $1;



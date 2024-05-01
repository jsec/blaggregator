-- name: CreateFollow :one
insert into feed_follows (id, user_id, feed_id, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning *;

-- name: DeleteFollow :exec
delete from feed_follows
where id = $1
and user_id = $2;

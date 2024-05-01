-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select * from feeds
order by id;

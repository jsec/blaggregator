-- +goose Up
alter table users
add api_key varchar(64) not null,
add constraint users_api_key unique (api_key);

-- +goose Down
alter table users
drop constraint users_api_key,
drop column api_key;

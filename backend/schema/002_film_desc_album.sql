-- +migrate Up
alter table films add column if not exists descr text;
alter table films add column if not exists album_id varchar(255);


-- +migrate Down
alter table films drop column if exists descr;
alter table films drop column if exists album_id;

-- +migrate Up

alter table device
  add column api_key character varying(32) not null default '';

-- +migrate Down

alter table device
  drop colum api_key;

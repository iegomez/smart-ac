-- +migrate Up

alter table datum
  add column created_at timestamp with time zone not null;

create index idx_datum_created_at on datum(created_at);

-- +migrate Down
drop index idx_datum_created_at;

alter table datum
  drop column created_at;

alter table datum
  add column created_at timestamp not null;

add index idx_datum_created_at on datum(created_at);
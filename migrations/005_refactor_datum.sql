-- +migrate Up

create index idx_user_username on "user" using gin(username gin_trgm_ops);
alter table datum
  drop column temperature,
  drop column carbon_monoxide,
  drop column air_humidity,
  drop column health_status,
  add column sensor_type character varying(20) not null default '',
  add column val double precision not null default 0.0,
  add column str_val text not null default '';

create index idx_datum_sensor_type on datum using gin (sensor_type gin_trgm_ops);

-- +migrate Down

drop index idx_datum_sensor_type;
drop index idx_user_username;
alter table datum
  drop column sensor_type,
  drop column val,
  drop column str_val,
  add column temperature double precision not null default 0.0,
  add column carbon_monoxide double precision not null default 0.0,
  add column air_humidity double precision not null default 0.0,
  add column health_status text not null default '';
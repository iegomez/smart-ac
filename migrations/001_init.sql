-- +migrate Up

create table device (
  id bigserial primary key,
	serial_number text not null default '',
  registered_at timestamp not null,
  firmware_version text not null default ''
);

create index idx_device_serial_number_trgm on device using gin (serial_number gin_trgm_ops);
create index idx_device_registered_at on device(registered_at);

create table datum (
  id bigserial primary key,
  device_id bigint references device on delete cascade,
	temperature double precision not null default 0.0,
	carbon_monoxide double precision not null default 0.0,
  air_humidity double precision not null default 0.0,
  health_status text not null default ''
);

-- +migrate Down
drop index idx_device_serial_number_trgm;
drop index idx_device_registered_at;
drop table datum;
drop table device;
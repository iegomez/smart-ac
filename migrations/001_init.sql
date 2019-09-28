-- +migrate Up

create table "user" (
	id bigserial primary key,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone not null,
	username character varying (100) not null,
	password_hash character varying (200) not null,
	session_ttl bigint not null,
	is_admin boolean not null
);

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

-- Create an initial admin (password: admin)
insert into "user" (
	created_at,
	updated_at,
	username,
	password_hash,
	session_ttl,
	is_admin
) values (
	now(),
	now(),
	'admin',
	'PBKDF2$sha512$1$l8zGKtxRESq3PA2kFhHRWA==$H3lGMxOt55wjwoc+myeOoABofJY9oDpldJa7fhqdjbh700V6FLPML75UmBOt9J5VFNjAL1AvqCozA1HJM0QVGA==',
	0,
	true
);

-- +migrate Down
drop index idx_device_serial_number_trgm;
drop index idx_device_registered_at;
drop table datum;
drop table device;
drop table "user";
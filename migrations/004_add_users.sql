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
	'$2a$14$08DXPq9ShIzDoRe76ZvPZeqfM7iR2RYFFighZhaPAUPGqfd.W7./y',
	0,
	true
);

-- +migrate Down

drop table "user";
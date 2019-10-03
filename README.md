# smart-ac

### Building

The project comes with a Makefile to ease building the application. Cleaning up and building both the backend and the frontend may be done like:
```
make clean build
```
This will generate the `smart-ac` binary with all static files embedded at the `build` dir.

### API definitions

You may find the api definitions and coumentation at https://iegomez.cl/api.

### Configuration

The app expect a `toml` file to gather its configuration. By defualt it'll look for `smart-ac.toml` at the same `dir` as the binary, but a path may be given with the `--conf` flag.  
The conf file should look like this:
```toml
[general]

  log_level = 4

[postgresql]

  # the postgres DB dsn
  dsn="postgres://smart_ac:password@localhost/smart_ac?sslmode=disable"
  # should the application run migrations on startup
  automigrate=true

[external_api]
  # ip:port to bind the (user facing) http server to (web-interface and REST / gRPC api)  
  bind="0.0.0.0:8080"

  # http server TLS certificate (optional)
  tls_cert=""

  # http server TLS key (optional)
  tls_key=""

  # JWT secret used for api authentication / authorization
  jwt_secret="your-jwt-secret"
```

### Testing

You may test the application given that you have a Postgres test DB `smart_ac_test`, owned by the user `smart_ac` with password `password`. The DB should also have the extension `pg_trgm` enabled, which may be created by a Postgres superuser like this:

```psql
psql$: \c smart_ac_test
psql$: create extension pg_trgm;
```

of course, the production DB must also have the `pg_trgm` extension enabled.
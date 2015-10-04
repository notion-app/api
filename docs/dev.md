
Information about developing the notion API.

# Installing

`go v1.5` and `gb` are required (google `gb go build` for the last one).

1. `make` will compile a binary in `bin/main`

2. `make run` will run it.

3. `make test` will run tests.

# Envvars

```
$ENV           (fails if not provided)
$PORT          (default: '8080')
$DATABASE_URL  (fails if not provided)
```

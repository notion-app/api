
Information about developing the notion API.

# Installing

You **need** Node 4.0. We use various ES6 functionality and Node4 is the first version to support it. No, we do not use Babel. Just upgrade to Node4.

1. `npm install` to install all dependencies.

2. `npm install -g bunyan` to globally install the bunyan logging framework. This is optional but necessary for well formatted logs (and `make` uses it).

3. `make tunnel` will set up SSH port forwarding on 8081 and 28015 for the RethinkDB cluster. You'll need SSH access into the cluster; ask Michael.

4. `make` will run the API server on 8080.

# Envvars

### Required

This thing is required due to a bug in knex right now.

```
export PGSSLMODE="require"
```

### Optional

```
$ENV           (default 'dev')
$PORT          (default '8080')
$DATABASE_URL  (default notion-dev db on heroku)
```

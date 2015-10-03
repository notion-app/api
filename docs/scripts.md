
There are a few scripts provided in `/scripts`

# knex_base

This is a template node script for writing knex scripts in should you need to run a run-off job.

# new_user

This posts a new user into an already-running notion api to test endpoints.

# reset_db

This drops all the tables in the database under $DATABASE_URL and recreates them.
More interestingly, this file defines the schema for our database tables. So its pretty important.
Just do be careful. Obviously.

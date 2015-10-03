# Testing

Our test cases are completely comprehensive. There are some unit tests, but
there are others which test service-level functionality. This means you will
need database access and things like personal Facebook auth tokens set.

I know this is bad testing practice, but for me at least it makes a lot of
sense because it completely verifies that everything at every level should
be working exactly as we expect when we deploy.

This also means tests can take a long time to run because many have to
connect to external services like facebook and the database.

That being said, the tests will never leave traces behind after running.
They may create objects in the database, but they will always delete them
after creation.

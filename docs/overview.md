
# System Overview

### API Server

The entire API server is a single Node repository (`src/`). This is hosted on heroku.

We have two heroku projects: `notion-dev.herokuapp.com` and `notion.herokuapp.com`. Each of these has their own postgres and redis instance.

### Persistent Storage

Our database is Postgres hosted by Heroku.

The URLs for these are located in `src/config.js`. Yeah. We're storing secrets in our repository.
Remind me to turn in my programming license to Steve Ballmer as soon as possible.

I'd suggest setting `$DATABASE_URL` to the 'dev' URL.

Why switch from Rethink? Rethink offered value in the real time updating, but I had the following concerns:

1. Validation logic. There will still be validation logic in the API, but the additional layer of safety
2. Defined schemas. Obvious.
3. Real time updating wasn't as valuable as I thought considering we are still using redis for caching and it has pubsub.
4. Hosting our own database is scary. 

### RTM

Real time updating of notes is provided through Redis PubSub, hosted on Heroku. The URLs for these are in, you guessed it, `src/config.js`.

### Flow

Basic CRUD API operations should be easy to implement and not worth iterating here.

The real time note updating is more interesting, because it is confounded by the following two problems:

1. Unlike Rethink, postgres does not have the built-in ability to notify clients on change.
2. Node isn't multithreaded.

Thus, the system I am considering is organized something like this:

1. User opens a notebook `nb_id`. This is a GET request which returns a websocket which is then connected to.
2. API loads the entire notebook content into redis and creates and subscribes to a pubsub channel with name `nb_id` and `class_id`.
3. User requests an update to some note with id `note_id` in the notebook.
4. API "synchronously" updates the note in redis, asynchronously updates the note in postgres, then publishes an update message to the redis channels created before
5. Elsewhere in the API, we receive the update through the redis channel and can kick off anything we need there at that point.

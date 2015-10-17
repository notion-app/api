
psql $DATABASE_URL -c "CREATE TABLE schools (
  id text PRIMARY KEY,
  name text NOT NULL,
  location text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE courses (
  id text PRIMARY KEY,
  school_id text REFERENCES schools (id) NOT NULL,
  name text NOT NULL,
  number text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE sections (
  id text PRIMARY KEY,
  course_id text REFERENCES courses (id) NOT NULL,
  crn text,
  professor text,
  year text NOT NULL,
  semester text NOT NULL,
  time text,
  verified boolean
)"

psql $DATABASE_URL -c "CREATE TABLE users (
  id text PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  verified boolean,
  auth_method text NOT NULL,

  school text REFERENCES schools (id) DEFAULT NULL,

  fb_user_id text,
  fb_auth_token text,
  fb_profile_pic text
)"

psql $DATABASE_URL -c "CREATE TABLE school_requests (
  id text PRIMARY KEY,
  requester_user_id text REFERENCES users (id) NOT NULL,
  name text NOT NULL,
  location text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE notebooks (
  id text PRIMARY KEY,
  section_id text,
  name text,
  owner text REFERENCES users (id),
  privacy text
)"

psql $DATABASE_URL -c "CREATE TABLE topics (
  id text PRIMARY KEY,
  notebook_id text REFERENCES notebooks (id)
)"

psql $DATABASE_URL -c "CREATE TABLE notes (
  id text PRIMARY KEY,
  topic_id text REFERENCES topics (id) NOT NULL,
  name text NOT NULL,
  owner text REFERENCES users (id) NOT NULL,
  endorsements integer DEFAULT 0,
  content text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE subscriptions (
  user_id text REFERENCES users (id) NOT NULL,
  notebook_id text REFERENCES notebooks (id) NOT NULL,
  PRIMARY KEY (user_id, notebook_id)
)"

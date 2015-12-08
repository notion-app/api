
psql $DATABASE_URL -c "CREATE TABLE schools (
  id text PRIMARY KEY,
  name text NOT NULL,
  location text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE users (
  id text PRIMARY KEY,
  username text,
  name text NOT NULL,
  email text NOT NULL,
  verified boolean,
  auth_method text NOT NULL,

  school text REFERENCES schools (id) ON DELETE RESTRICT DEFAULT NULL,

  fb_user_id text,
  fb_auth_token text,
  fb_profile_pic text
)"

psql $DATABASE_URL -c "CREATE TABLE courses (
  id text PRIMARY KEY,
  school_id text REFERENCES schools (id) ON DELETE CASCADE NOT NULL,
  owner text REFERENCES users (id),
  name text NOT NULL,
  number text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE notebooks (
  id text PRIMARY KEY,
  name text,
  owner text REFERENCES users (id) ON DELETE CASCADE DEFAULT NULL,
  privacy text,
  created_at timestamptz,
  updated_at timestamptz
)"

psql $DATABASE_URL -c "CREATE TABLE sections (
  id text PRIMARY KEY,
  course_id text REFERENCES courses (id) ON DELETE CASCADE NOT NULL,
  notebook_id text REFERENCES notebooks (id) ON DELETE CASCADE NOT NULL,
  crn text,
  professor text,
  year text NOT NULL,
  semester text NOT NULL,
  time text,
  verified boolean
)"

psql $DATABASE_URL -c "CREATE TABLE school_requests (
  id text PRIMARY KEY,
  requester_user_id text REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  name text NOT NULL,
  location text NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE topics (
  id text PRIMARY KEY,
  notebook_id text REFERENCES notebooks (id) ON DELETE CASCADE NOT NULL
)"

psql $DATABASE_URL -c "CREATE TABLE notes (
  id text PRIMARY KEY,
  topic_id text REFERENCES topics (id) ON DELETE RESTRICT NOT NULL,
  title text,
  owner text REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  content text NOT NULL,
  created_at timestamptz,
  updated_at timestamptz
)"

psql $DATABASE_URL -c "CREATE TABLE subscriptions (
  user_id text REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  notebook_id text REFERENCES notebooks (id) ON DELETE CASCADE NOT NULL,
  name text,
  PRIMARY KEY (user_id, notebook_id)
)"

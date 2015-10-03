
psql $DATABASE_URL -c "CREATE TABLE schools (
  id text,
  name text,
  location text,

  PRIMARY KEY (id)
)"

# We manually insert Purdue here as a debug school so that way it has the same ID every time
psql $DATABASE_URL -c "INSERT INTO schools VALUES (
  'hROIgEBrcM',
  'Purdue University',
  'West Lafayette, IN'
)"

psql $DATABASE_URL -c "CREATE TABLE users (
  id text NOT NULL,
  name text NOT NULL,
  role text NOT NULL,
  verified boolean,
  auth_method text NOT NULL,

  fb_user_id text,
  fb_auth_token text,
  fb_expires_in text,
  fb_profile_pic text,

  PRIMARY KEY (id)
)"

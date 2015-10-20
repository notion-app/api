
sh ./scripts/drop_tables.sh
sh ./scripts/create_tables.sh

echo 'Populate Schools'
psql $DATABASE_URL -c "INSERT INTO schools VALUES
  (
    'School1',
    'My Fake University',
    'Hometown USA'
  ),
  (
    'School2',
    'Another Fake College',
    'Awesome Town, CA'
  )"

echo 'Populate Courses'
psql $DATABASE_URL -c "INSERT INTO courses VALUES
  (
    'Course1',
    'School1',
    'Programming in Java',
    'CS 180'
  ),
  (
    'Course2',
    'School1',
    'Programming in C',
    'CS 240'
  ),
  (
    'Course3',
    'School2',
    'Advanced English',
    'ENGL 410'
  )"

echo 'Populate Sections'
psql $DATABASE_URL -c "INSERT INTO sections VALUES
  (
    'Section11',
    'Course1',
    '111',
    'Dr. Feel Good',
    '2015',
    'Fall',
    '3:30 PM',
    false
  ),
  (
    'Section12',
    'Course1',
    '112',
    'Richard Stallman',
    '2014',
    'Fall',
    '1:30 PM',
    false
  ),
  (
    'Section21',
    'Course2',
    '111',
    'Dr. Michael Hockerman',
    '2015',
    'Spring',
    '12:30 PM',
    false
  )"

echo 'Populate Users'
psql $DATABASE_URL -c "INSERT INTO users VALUES
  (
    'User1',
    'Bob Yelpington',
    'bob@website.com',
    false,
    'facebook',
    NULL,
    '12345',
    'abcde',
    'http://google.com'
  ),
  (
    'User2',
    'Nichole Clayton',
    'me@memememe.com',
    false,
    'facebook',
    'School1',
    'FbUser1',
    'FbAuthToken1',
    'http://profile-pic.com'
  )"

echo 'Populate Notebooks'
psql $DATABASE_URL -c "INSERT INTO notebooks VALUES
  (
    'Notebook1',
    'Section1',
    NULL,
    NULL,
    'section'
  ),
  (
    'Notebook2',
    'Section2',
    NULL,
    NULL,
    'section'
  )"

echo 'Populate Topics'
psql $DATABASE_URL -c "INSERT INTO topics VALUES
  (
    'Topic1',
    'Notebook1'
  ),
  (
    'Topic2',
    'Notebook1'
  ),
  (
    'Topic3',
    'Notebook2'
  )"

echo 'Populate Notes'
psql $DATABASE_URL -c "INSERT INTO notes VALUES
  (
    'Note1',
    'Topic1',
    NULL,
    'User1',
    0,
    'This is an awesome note that i wrote'
  ),
  (
    'Note2',
    'Topic2',
    'My Title',
    'User1',
    1,
    'Ive got another note here, its pretty cool'
  ),
  (
    'Note3',
    'Topic1',
    'Another Title',
    'User2',
    0,
    'Wow these notes are amazing arent they?'
  )"

echo 'Populate Subscriptions'
psql $DATABASE_URL -c "INSERT INTO subscriptions VALUES
  (
    'User1',
    'Notebook1'
  ),
  (
    'User1',
    'Notebook2'
  ),
  (
    'User2',
    'Notebook1'
  )"
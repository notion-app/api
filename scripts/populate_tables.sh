# We manually insert Purdue here as a debug school so that way it has the same ID every time
while getopts ":d" opt; do
  case $opt in
    d)
      echo "Dropping Tables" >&2
      sh ./drop_tables.sh
      sh ./create_tables.sh
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done


echo 'Populate schools'
psql $DATABASE_URL -c "INSERT INTO schools VALUES (
  'hROIgEBrcM',
  'Purdue University',
  'West Lafayette, IN'
)"

echo 'Populate courses'
psql $DATABASE_URL -c "INSERT INTO courses VALUES
  (
    'gAXkjhACll',
    'hROIgEBrcM',
    'Programming in Java',
    'CS 180'
  ),
  (
    'WHZGSCjYCg',
    'hROIgEBrcM',
    'Programming in C',
    'CS 240'
  )"

echo 'Populate sections'
psql $DATABASE_URL -c "INSERT INTO sections VALUES
  (
    'dgGyXeUJbm',
      'gAXkjhACll',
      '184825',
      'Dr. Feel Good',
      '2015',
      'Fall',
      '3:30 PM',
      false
    ),
    (
      'CMrlxrBeNV',
      'WHZGSCjYCg',
      '112349',
      'Richard Stallman',
      '2015',
      'Fall',
      '1:30 PM',
      false
    )"

echo 'Populate notebooks'
psql $DATABASE_URL -c "INSERT INTO notebooks VALUES
    (
      'VFZNcusCbQ',
      'dgGyXeUJbm',
      'CS 180 Notebook',
      null,
      'section'
    ),
    (
      'hwnwNpCbJJ',
      'CMrlxrBeNV',
      'CS 240 Notebook',
      null,
      'section'
      )"

echo 'Populate topics'
psql $DATABASE_URL -c "INSERT INTO topics VALUES
    (
      'DrUrxIjQwu',
      'VFZNcusCbQ'
    ),
    (
      'pxYjbBFgvM',
      'hwnwNpCbJJ'
    )"

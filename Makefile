
build:
	gb build

run: build
	bin/app

internaltest:
	gb test

remakedb:
	./scripts/debug_db.sh

logs:
	heroku logs --app notion-api-dev -p web.1 --tail

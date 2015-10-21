
build:
	gb build

run: build
	bin/app

internaltest:
	gb test

remakedb:
	./scripts/debug_db.sh

logs:
	heroku logs -p web.1 --tail

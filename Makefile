
build:
	gb build

run: build
	bin/app

test:
	gb test

logs:
	heroku logs -p web.1 --tail

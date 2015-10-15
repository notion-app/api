
build:
	gb build

run: build
	bin/app

test:
	gb test

hklogs:
	heroku logs -p web.1 --tail

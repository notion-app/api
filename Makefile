
build:
	go build

run:
	go run main.go

deps:
	go get -u github.com/labstack/echo
	go get -u github.com/labstack/echo/middleware
	go get -u github.com/Sirupsen/logrus

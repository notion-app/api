package routes

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"gitlab.com/notionapp/api/config"
	"gitlab.com/notionapp/api/log"
)

var (
	e = echo.New()
)

// Init registers our routes into echo
func Init() {
	log.Info("Initializing routes")
	middleware()
	v1Routes()
	e.Run(config.WebPort())
}

func middleware() {
	e.Use(mw.Recover())
}

func v1Routes() {
	e.Get("/status", Status)
}

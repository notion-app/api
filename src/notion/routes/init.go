package routes

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"notion/config"
	"notion/log"
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
	e.Get("/v1/status", Status)
	e.Get("/v1/user/:user_id", GetUser)
}

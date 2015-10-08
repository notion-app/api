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
	v1Group := e.Group("/v1")
	v1Group.Get("/status", Status)
	v1Group.Post("/login", Login)
	v1Group.Get("/school", GetSchools)
	v1Group.Post("/school/request", PostSchoolRequest)
	// e.Get("/v1/user/:user_id", GetUser)

}

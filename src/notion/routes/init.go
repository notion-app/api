package routes

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"notion/config"
	"notion/log"
	mmw "notion/middleware"
)

var (
	e = echo.New()
)

// Init registers our routes into echo
func Init() {
	log.Info("Initializing routes")
	middleware()
	v1Routes()
	log.Info("Serving API on port %v", config.WebPort())
	e.Run(config.WebPort())
}

func middleware() {
	e.Use(mw.Recover())
	e.Use(mmw.AccessControl())
	e.Use(mmw.ParseBody())
}

func v1Routes() {
	v1Group := e.Group("/v1")

	// Unauthenticated endpoints
	// status.go
	v1Group.Get("/status", Status)
	// login.go
	v1Group.Post("/login", Login)
	// schools.go
	v1Group.Get("/school", GetSchools)
	v1Group.Get("/school/:school_id/course", GetCoursesForSchool)
	v1Group.Get("/school/:school_id/course/:course_id/section", GetSectionsForCourse)

	// Authenticated endpoints
	authV1Group := v1Group.Group("")
	authV1Group.Use(mmw.TokenCheck())
	// schools.go
	authV1Group.Post("/school/request", PostSchoolRequest)
	// users.go
	authV1Group.Get("/user/:user_id", GetUser)
	authV1Group.Get("/user/:user_id/subscription", GetUsersSubscriptions)
	authV1Group.Post("/user/:user_id/subscription", CreateUserSubscription)
	authV1Group.Delete("/user/:user_id/subscription", RemoveUserSubscription)
	authV1Group.Put("/user/:user_id/school", SetUserSchool)
	authV1Group.Options("/user/:user_id/school", SetUserSchool)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"notion/config"
	"notion/log"
	mw "notion/middleware"
)

var (
	g = gin.New()
)

// Init registers our routes into Gin
func Init() {
	log.Info("Initializing routes")
	gin.SetMode(gin.ReleaseMode)
	middleware()
	v1Routes()
	log.Info("Serving API on port %v", config.WebPort())
	g.Run(config.WebPort())
}

func middleware() {
	g.Use(mw.Logger)
	g.Use(mw.AccessControl)
}

func v1Routes() {
	v1Group := g.Group("/v1")

	// Unauthed status.go
	v1Group.GET("/status", Status)

	// Unauthed login.go
	v1Group.POST("/login", Login)

	// Unauthed schools.go
	// v1Group.GET("/school", GetSchools)
	// v1Group.GET("/school/:school_id/course", GetCoursesForSchool)
	// v1Group.GET("/school/:school_id/course/:course_id/section", GetSectionsForCourse)
	//
	// // Authenticated endpoints
	// authV1Group := v1Group.Group("")
	//
	// // notebook.go
	// authV1Group.GET("/notebook/:notebook_id/topic", GetNotebookNotes)
	// // schools.go
	// authV1Group.POST("/school/request", PostSchoolRequest)
	// // users.go
	// authV1Group.GET("/user/:user_id", GetUser)
	// authV1Group.GET("/user/:user_id/subscription", GetUsersSubscriptions)
	// authV1Group.POST("/user/:user_id/subscription", CreateUserSubscription)
	// authV1Group.PUT("/user/:user_id/subscription", ModifyUserSubscription)
	// authV1Group.DELETE("/user/:user_id/subscription", RemoveUserSubscription)
	// authV1Group.PUT("/user/:user_id/school", SetUserSchool)
}

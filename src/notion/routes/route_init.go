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
	g.Use(mw.Error)
}

func v1Routes() {
	v1 := g.Group("/v1")
	v1.GET("/status", Status)
	v1.POST("/login", Login)

	v1a := v1.Group("", mw.AuthCheck)
	v1a.GET("/school", GetAllSchools)

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

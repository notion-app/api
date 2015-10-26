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
	v1a.GET("/school/:school_id/course", GetCoursesForSchool)
	v1a.GET("/school/:school_id/course/:course_id/section", GetSectionsForCourse)
	v1a.POST("/school/request", PostSchoolRequest)

	// v1a.GET("/notebook/:notebook_id/topic", GetNotebookNotes)

	v1a.GET("/user/:user_id", GetUser)
	v1a.GET("/user/:user_id/subscription", GetUsersSubscriptions)
	v1a.POST("/user/:user_id/subscription", CreateUserSubscription)
	v1a.PUT("/user/:user_id/subscription", ModifyUserSubscription)
	v1a.PUT("/user/:user_id/school", SetUserSchool)
	v1a.DELETE("/user/:user_id/subscription", RemoveUserSubscription)

}

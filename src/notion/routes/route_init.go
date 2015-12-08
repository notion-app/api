package routes

import (
	"github.com/gin-gonic/gin"
	"notion/config"
	"notion/log"
	mw "notion/middleware"
	v1 "notion/routes/v1"
	// v2 "notion/routes/v2"
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
	v2Routes()
	log.Info("Serving API on port %v", config.WebPort())
	g.Run(config.WebPort())
}

func middleware() {
	g.Use(mw.Logger)
	g.Use(mw.RateLimit)
	g.Use(mw.AccessControl)
	g.Use(mw.Error)
	g.Use(mw.AcceptOptions)
}

func v1Routes() {
	v1g := g.Group("/v1")

	v1g.GET("/status", v1.Status)
	v1g.POST("/login", v1.Login)

	v1a := v1g.Group("", mw.AuthCheck)

	v1g.GET("/school", v1.GetAllSchools)
	v1g.GET("/school/:school_id", v1.GetSingleSchool)
	v1g.GET("/school/:school_id/course", v1.GetCoursesForSchool)
	v1g.GET("/school/:school_id/course/:course_id/section", v1.GetSectionsForCourse)
	v1a.POST("/school/request", v1.PostSchoolRequest)

	v1a.GET("/notebook/:notebook_id/topic", v1.GetNotebookNotes)

	v1a.GET("/user/:user_id", v1.GetUser)
	v1a.PUT("/user/:user_id/email/:email", v1.SetUserEmail)
	v1a.PUT("/user/:user_id/username/:username", v1.SetUserUsername)
	v1a.GET("/user/:user_id/subscription", v1.GetUsersSubscriptions)
	v1a.POST("/user/:user_id/subscription", v1.CreateUserSubscription)
	v1a.PUT("/user/:user_id/subscription", v1.ModifyUserSubscription)
	v1a.PUT("/user/:user_id/school", v1.SetUserSchool)
	v1a.DELETE("/user/:user_id/subscription/:notebook_id", v1.RemoveUserSubscription)

	v1a.GET("/notebook/:notebook_id/note/:note_id", v1.GetSingleNote)
	v1a.POST("/notebook/:notebook_id/note", v1.CreateNote)
	v1a.PUT("/notebook/:notebook_id/note/:note_id", v1.ModifyNote)
	v1a.DELETE("/notebook/:notebook_id/note/:note_id", v1.DeleteNote)
	v1a.POST("/notebook/:notebook_id/note/:note_id/change", v1.PostNoteChange)

	v1g.GET("/echo", v1.EchoWebsocket)
	v1a.GET("/note/:note_id/ws", v1.OpenWebsocket)
}

func v2Routes() {
	v2g := g.Group("/v2")
	v2g.POST("/login", v1.Login)
	// v2a := v2g.
}

package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"keeper-crud/controller"
	"os"
)

func NewUsersRouter(baseRouter *gin.RouterGroup) *gin.RouterGroup {
	store := cookie.NewStore([]byte(os.Getenv("KEEPER_SESSIONS_SECRET")))
	baseRouter.Use(sessions.Sessions("users_sessions", store))
	return baseRouter.Group("/users")
}

func SetupUsersRouter(baseRouter *gin.RouterGroup, usersController *controller.UsersController) {
	usersRouter := NewUsersRouter(baseRouter)
	usersRouter.POST("/signup", usersController.Signup)
	usersRouter.POST("/signin", usersController.Signin)
}

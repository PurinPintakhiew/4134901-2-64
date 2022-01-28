package routers

import (
	"purin/lab8/controllers"

	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine){
	r.POST("/user/:id",controllers.AddUser)
	r.DELETE("/users",controllers.DelAllUser)
}
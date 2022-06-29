package routers

import (
	"github.com/gin-gonic/gin"

	"server/middleware/jwt"
	v1 "server/routers/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.CORSMiddleware())
	{
		apiv1.POST("/pomos/list", v1.GetPomos)
		apiv1.POST("/pomos/create", v1.CreatePomo)
		apiv1.POST("/pomos/count", v1.CountPomos)

		apiv1.POST("/todos/list", v1.GetTodos)
		apiv1.POST("/todos/create", v1.CreateTodo)
		apiv1.POST("/todos/delete", v1.DeleteTodo)
		apiv1.POST("/todos/update", v1.UpdateTodo)
		// apiv1.POST("/todos/count", v1.CountTodos)
	}
	return r
}

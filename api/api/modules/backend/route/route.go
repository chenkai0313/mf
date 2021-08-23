package route

import (
	"github.com/gin-gonic/gin"

	"api/api/modules/backend/controller"
	"api/api/modules/backend/middleware"
)

func Route(r *gin.Engine) {
	authorized := r.Group("/backend")
	authorized.Use(middleware.BackendMiddle())
	{
		authorized.POST("/test", controller.Test)
	}
}

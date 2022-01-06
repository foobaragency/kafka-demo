package services

import (
	"github.com/gin-gonic/gin"
)

func CreateHandlers(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/:name", CreateProductHandler)
	}
}

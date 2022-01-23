package services

import (
	"github.com/gin-gonic/gin"
)

func CreateHandlers(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		products := v1.Group("products")
		{
			products.POST("/", CreateProductHandler)
			products.PATCH("/:id/stock/:quantity", RefillStockHandler)
		}
	}
}

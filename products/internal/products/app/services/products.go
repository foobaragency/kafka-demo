package services

import (
	"net/http"

	"github.com/foobaragency/kafka-demo/products/internal/products/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProductHandler(c *gin.Context) {
	name := c.Param("name")

	id := uuid.NewString()

	p, err := domain.CreateProduct(&domain.Product{
		ID:   id,
		Name: name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusAccepted, p)
}

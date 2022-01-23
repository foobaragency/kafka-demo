package services

import (
	"net/http"
	"strconv"

	"github.com/foobaragency/kafka-demo/products/internal/products/domain"
	"github.com/gin-gonic/gin"
)

func CreateProductHandler(c *gin.Context) {
	var product domain.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := domain.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, product)
}

func RefillStockHandler(c *gin.Context) {
	id := c.Param("id")
	quantity := c.Param("quantity")

	quantityInt, err := strconv.ParseInt(quantity, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product := &domain.Product{
		ID:    id,
		Stock: quantityInt,
	}

	if err := domain.RefillStock(product); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, product)
}

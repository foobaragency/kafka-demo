package domain

import (
	"github.com/foobaragency/kafka-demo/products/pkg/events"
)

type Product struct {
	events.EventData
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Stock       int64  `json:"stock"`
}

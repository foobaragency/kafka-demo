package domain

import (
	"github.com/foobaragency/kafka-demo/products/pkg/events"
)

type Product struct {
	events.EventData
	ID   string `json:"id"`
	Name string `json:"name"`
}

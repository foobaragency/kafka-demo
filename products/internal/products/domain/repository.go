package domain

import (
	"context"
	"time"

	"github.com/foobaragency/kafka-demo/products/pkg/events"
)

func CreateProduct(p *Product) (*Product, error) {
	metadata := &events.EventMetadata{
		EventType:        "kafka_demo.CreateProduct",
		EventSource:      "productAPI/:name",
		EventKey:         p.ID,
		EventDestination: "products",
		EventTime:        time.Now(),
	}

	err := p.CreateEvent(context.Background(), metadata)

	if err != nil {
		return nil, err
	}

	err = p.Event.SetEventData(p)

	if err != nil {
		return nil, err
	}

	err = p.Event.Send()

	if err != nil {
		return nil, err
	}

	return p, nil
}

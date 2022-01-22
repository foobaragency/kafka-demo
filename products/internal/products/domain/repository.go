package domain

import (
	"context"
	"time"

	"github.com/foobaragency/kafka-demo/products/pkg/events"
)

func CreateProduct(p *Product) error {
	metadata := &events.EventMetadata{
		EventType:        "kafka_demo.CreateProduct",
		EventSource:      "productsAPI/v1/products",
		EventKey:         p.ID,
		EventDestination: "products",
		EventTime:        time.Now(),
	}

	err := p.CreateEvent(context.Background(), metadata)

	if err != nil {
		return err
	}

	err = p.Event.SetEventData(p)

	if err != nil {
		return err
	}

	err = p.Event.Send()

	if err != nil {
		return err
	}

	return nil
}

func RefillStock(p *Product) error {
	metadata := &events.EventMetadata{
		EventType:        "kafka_demo.RefillStock",
		EventSource:      "productsAPI/v1/products/:id/stock/:quantity",
		EventKey:         p.ID,
		EventDestination: "products",
		EventTime:        time.Now(),
	}

	err := p.CreateEvent(context.Background(), metadata)

	if err != nil {
		return err
	}

	err = p.Event.SetEventData(p)

	if err != nil {
		return err
	}

	err = p.Event.Send()

	if err != nil {
		return err
	}

	return nil
}

package events

import (
	"context"
	"time"

	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Event struct {
	cloudevents.Event
	destination string
	eventKey    string
	client      cloudevents.Client
	sender      *kafka_sarama.Sender
	ctx         context.Context
	sendResult  protocol.Result
}

type EventMetadata struct {
	EventType        string
	EventSource      string
	EventKey         string
	EventDestination string
	EventTime        time.Time
}

type EventData struct {
	EventID     string    `json:"ce_id"`
	EventType   string    `json:"ce_type"`
	EventSource string    `json:"ce_source"`
	EventTime   time.Time `json:"ce_time"`
	Event       *Event    `json:"-"`
}

func (e *EventData) CreateEvent(ctx context.Context, metadata *EventMetadata) error {
	var ev *Event
	ev, err := NewEvent(ctx, metadata)

	if err != nil {
		return err
	}

	e.Event = ev

	e.EventID = ev.ID()
	e.EventType = ev.Type()
	e.EventTime = ev.Time()
	e.EventSource = ev.Source()

	return nil
}

func (e *Event) SetEventID(id string) {
	e.SetID(id)
}

func (e *Event) SetContext(ctx context.Context) {
	e.ctx = ctx
}

func (e *Event) Destination() string {
	return e.destination
}

func (e *Event) SetDestination(destination string) {
	e.destination = destination
}

func (e *Event) Key() string {
	return e.eventKey
}

func (e *Event) SetKey(key string) {
	e.eventKey = key
}

func (e *Event) SetEventData(data interface{}) error {
	return e.SetData(cloudevents.ApplicationJSON, data)
}

func NewEvent(ctx context.Context, metadata *EventMetadata) (*Event, error) {
	var e *Event

	eventId := uuid.NewString()

	log.Info().Msg(eventId)

	e = &Event{}

	e.SetSpecVersion(cloudevents.VersionV1)

	e.SetContext(ctx)
	e.SetID(eventId)
	e.SetType(metadata.EventType)
	e.SetSource(metadata.EventSource)
	e.SetTime(metadata.EventTime)
	e.SetKey(metadata.EventKey)
	e.SetDestination(metadata.EventDestination)

	err := e.Validate()

	if err != nil {
		return nil, err
	}

	err = e.BuildSender()

	if err != nil {
		return nil, err
	}

	return e, nil
}

package events

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func KafkaEventError(event *Event, err error, msg string) {
	log.Error().Dict(
		"information",
		zerolog.Dict().
			Str("resource", "kafka").
			Str("topic", event.Destination()).
			Str("eventID", event.ID()).
			Str("eventType", event.Type()).
			Str("eventSource", event.Source()).
			Err(err),
	).Msg(msg)
}

func KafkaEventDelivered(event *Event) {
	log.Info().Dict(
		"information",
		zerolog.Dict().
			Str("resource", "kafka").
			Str("topic", event.Destination()).
			Str("eventID", event.ID()).
			Str("eventType", event.Type()).
			Str("eventSource", event.Source()),
	).Msg("event delivered!")
}

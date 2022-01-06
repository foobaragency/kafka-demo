package events

import (
	"crypto/tls"
	"errors"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func (e *Event) BuildSender() error {
	brokerList := os.Getenv("KAFKA_BROKERS")
	kafkaUser := os.Getenv("KAFKA_USER")
	kafkaPassword := os.Getenv("KAFKA_PASSWORD")
	brokers := strings.Split(brokerList, ",")

	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Net.SASL.Enable = true
	saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	saramaConfig.Net.SASL.User = kafkaUser
	saramaConfig.Net.SASL.Password = kafkaPassword
	saramaConfig.Net.SASL.Handshake = true
	saramaConfig.Net.TLS.Enable = true
	saramaConfig.Net.TLS.Config = &tls.Config{}

	sender, err := kafka_sarama.NewSender(brokers, saramaConfig, e.Destination())
	if err != nil {
		return err
	}

	e.sender = sender

	client, err := cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return err
	}

	e.client = client

	return nil
}

func (e *Event) Send() error {
	if result := e.client.Send(
		kafka_sarama.WithMessageKey(e.ctx, sarama.StringEncoder(e.Key())),
		e.Clone(),
	); cloudevents.IsUndelivered(result) {
		err := errors.New("failed to deliver event")
		KafkaEventError(e, err, err.Error())
		return err
	} else {
		KafkaEventDelivered(e)
	}
	return nil
}

func (e *Event) IsACK() bool {
	return cloudevents.IsACK(e.sendResult)
}

func (e *Event) Close() {
	defer e.sender.Close(e.ctx)
}

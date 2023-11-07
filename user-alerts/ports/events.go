package ports

import (
	"context"
	"encoding/json"
	"log"
	"user-alerts/app"
	"user-alerts/app/event"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

type Event struct {
	app        app.Application
	router     *message.Router
	subscriber *kafka.Subscriber
}

func NewEvents(app app.Application) *Event {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	e := &Event{app, router, subscriber}

	return e
}

func (e *Event) Start() {
	e.registerRoutes()
	e.router.AddPlugin(plugin.SignalsHandler)
	if err := e.router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func (e *Event) registerRoutes() {
	e.router.AddNoPublisherHandler(
		"user_created_handler",
		"users",
		e.subscriber,
		func(msg *message.Message) error {
			// Log the received message
			log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

			// Unmarshal the message payload into the UserCreated event
			var userCreatedEvent event.UserCreated
			if err := json.Unmarshal(msg.Payload, &userCreatedEvent); err != nil {
				log.Printf("error unmarshaling message payload: %v", err)
				return err
			}

			// Handle the event
			if err := e.app.Events.UserCreated.Handle(context.Background(), &userCreatedEvent); err != nil {
				log.Printf("error during event handling: %v", err)
				return err
			}

			return nil
		},
	)
}

package subscribe

import (
	"context"
	"encoding/json"
	"service2/connection"
	"service2/events/publish"
	"service2/proto/mail"
	"time"

	"go.opencensus.io/trace"

	"github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
)

// Subscriber allows subscription.
type Subscriber struct {
	pub       publish.Publisher
	closeChan chan struct{}
}

// New return new subscriber
func New(pub publish.Publisher) *Subscriber {
	return &Subscriber{pub: pub}
}

// Title returns events title.
func (sub *Subscriber) Title() string {
	return "Subscriber"
}

// Start starts event module
func (sub *Subscriber) Start() {
	sub.closeChan = make(chan struct{})

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true               // nolint
	config.Group.Return.Notifications = true           // nolint
	config.Group.Mode = cluster.ConsumerModePartitions // nolint

	// init consumer
	consumer, err := cluster.NewConsumer([]string{"localhost:9092"}, "service", []string{"service1"}, config)
	if err != nil {
		logrus.WithError(err).Fatal("new consumer")
	}
	defer func() {
		err = consumer.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			logrus.WithError(err).Error("consume error")
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			logrus.Infof("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume partitions
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				return
			}

			// start a separate goroutine to consume messages
			go sub.consumePartition(consumer, part)

		case <-sub.closeChan:
			return
		}
	}
}

func (sub *Subscriber) consumePartition(consumer *cluster.Consumer, pc cluster.PartitionConsumer) {
	for msg := range pc.Messages() {
		func() {
			consumer.MarkOffset(msg, "")

			txMail := &mail.MailTransaction{}
			err := txMail.Unmarshal(msg.Value)
			if err != nil {
				logrus.Errorf("Consume: %s. Proto unmarshal error: %s", txMail, err)
				return
			}

			err = txMail.Validate()
			if err != nil {
				logrus.Errorf("Consume: %+v\n. Validate error: %s", txMail, err)
				return
			}

			sub.service(txMail)
		}()
	}
}

func (sub *Subscriber) service(txMail *mail.MailTransaction) {

	///////////////////////////////// Trace /////////////////////////////////////////////
	// getting trace info from transaction
	var spanContext trace.SpanContext
	err := json.Unmarshal([]byte(txMail.Trace), &spanContext)
	if err != nil {
		logrus.Errorf("consumePartition, json.Unmarshal: %v\n", err)
		return
	}

	ctx, span := trace.StartSpanWithRemoteParent(
		context.Background(), "service2.consumePartition", spanContext)
	defer span.End()
	////////////////////////////////////////////////////////////////////////////////////

	con, err := connection.GetConnection(ctx)
	if err != nil {
		logrus.Errorf("service: error: %s", err)
	}
	txMail.ConnectionId = con.ConnectionID

	time.Sleep(2 * time.Second)

	spanContextJson, err := json.Marshal(span.SpanContext())
	if err != nil {
		return
	}
	txMail.Trace = string(spanContextJson)
	err = sub.pub.Publish(txMail)
	if err != nil {
		logrus.Errorf("service: error: %s", err)
	}
}

// Stop stops event module
func (sub *Subscriber) Stop() {
	close(sub.closeChan)
}

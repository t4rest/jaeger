package publish

import (
	"fmt"
	"service1/proto/mail"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// Publisher interface
type Publisher interface {
	Publish(evn *mail.MailTransaction) error
}

type kafkaPub struct {
	producer sarama.SyncProducer
}

// New creates new kafka connection
func New() (*kafkaPub, error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}

	return &kafkaPub{producer: producer}, nil
}

// Publish publish event
func (kfk *kafkaPub) Publish(txMail *mail.MailTransaction) error {
	data, err := txMail.Marshal()
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}
	_, _, err = kfk.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "service1",
		Value: sarama.ByteEncoder(data),
	})

	return err
}

// Close close connection
func (kfk *kafkaPub) Close() {

	err := kfk.producer.Close()
	if err != nil {
		logrus.Error(err)
	}
}

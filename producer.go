package stenographer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type producer struct {
	syncProducer sarama.SyncProducer
}

func newProducer(brokerAddress []string) *producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll

	syncProducer, err := sarama.NewSyncProducer(brokerAddress, config)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return &producer{
		syncProducer: syncProducer,
	}
}

func (p *producer) sendMessage(topic string, data []byte) (partition int32, offset int64, err error) {
	var saramaByte sarama.ByteEncoder
	saramaByte = data

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: saramaByte,
	}

	partition, offset, err = p.syncProducer.SendMessage(msg)
	return
}

func (p *producer) close() error {
	return p.syncProducer.Close()
}

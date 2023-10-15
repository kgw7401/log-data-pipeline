package kafka

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/jsonschema"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (k *KafkaProducer, err error) {
	conf := ReadConfig()

	p, err := kafka.NewProducer(&conf)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: p,
	}, nil
}

func (k *KafkaProducer) Serialize(topic string, data map[string]interface{}) ([]byte, error) {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig("http://localhost:8081"))

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		os.Exit(1)
	}

	serdeConfig := jsonschema.NewSerializerConfig()
	serdeConfig.AutoRegisterSchemas = false
	serdeConfig.UseLatestVersion = true

	ser, err := jsonschema.NewSerializer(client, serde.ValueSerde, serdeConfig)
	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		os.Exit(1)
	}

	payload, err := ser.Serialize(topic, &data)

	if err != nil {
		fmt.Printf("Failed to serialize data: %s\n", err)
		os.Exit(1)
	}

	return payload, err
}

// func (k *KafkaProducer) Serialize(topic string, data map[string]interface{}) ([]byte, error) {
// 	schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://localhost:8081")
// 	schema, err := schemaRegistryClient.GetLatestSchema(topic + "-value")
// 	if err != nil {
// 		panic(fmt.Sprintf("Error getting the schema %s", err))
// 	}

// 	err = schema.JsonSchema().Validate(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	d, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaIDBytes := make([]byte, 4)
// 	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))

// 	var recordValue []byte
// 	recordValue = append(recordValue, byte(0))
// 	recordValue = append(recordValue, schemaIDBytes...)
// 	recordValue = append(recordValue, d...)

// 	return recordValue, nil
// }

func (k *KafkaProducer) Produce(topic string, data []byte) {
	deliveryChan := make(chan kafka.Event)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, deliveryChan)
	if err != nil {
		fmt.Printf("Produce failed: %v\n", err)
		os.Exit(1)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

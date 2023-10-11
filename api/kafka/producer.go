package kafka

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/riferrei/srclient"
	"k8s.io/apimachinery/pkg/util/json"
)

func ProduceData(topic string, data map[string]interface{}) {

	conf := ReadConfig()

	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	defer p.Close()

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Failed to marshal data: %s\n", err)
		os.Exit(1)
	}

	schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://localhost:8081")
	schema, err := schemaRegistryClient.GetLatestSchema(topic + "-value")
	if err != nil {
		panic(fmt.Sprintf("Error getting the schema %s", err))
	}

	err = schema.JsonSchema().Validate(data)
	if err != nil {
		panic(fmt.Sprintf("Error validating the data %s", err))
	}
	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))

	var recordValue []byte
	recordValue = append(recordValue, byte(0))
	recordValue = append(recordValue, schemaIDBytes...)
	recordValue = append(recordValue, d...)
	// client, err := schemaregistry.NewClient(schemaregistry.NewConfig("http://localhost:8081"))

	// if err != nil {
	// 	fmt.Printf("Failed to create schema registry client: %s\n", err)
	// 	os.Exit(1)
	// }

	// serdeConfig := jsonschema.NewSerializerConfig()
	// serdeConfig.AutoRegisterSchemas = false
	// serdeConfig.UseLatestVersion = true
	// serdeConfig.EnableValidation = true

	// ser, err := jsonschema.NewSerializer(client, serde.ValueSerde, serdeConfig)

	// if err != nil {
	// 	fmt.Printf("Failed to create serializer: %s\n", err)
	// 	os.Exit(1)
	// }

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)
	// payload, err := ser.Serialize(topic, &d)

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed to serialize payload: %s\n", err)
		os.Exit(1)
	}
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          recordValue,
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

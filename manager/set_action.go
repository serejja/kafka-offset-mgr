package manager

import (
	"errors"
	kafka "github.com/serejja/kafka-client"
	"github.com/urfave/cli"
	"strings"
)

// ErrBrokersRequired happens when brokers flag is not set.
var ErrBrokersRequired = errors.New("At least one broker address is required.")

// ErrGroupRequired happens when group flag is not set.
var ErrGroupRequired = errors.New("Consumer group name is required.")

// ErrTopicRequired happens when topic flag is not set.
var ErrTopicRequired = errors.New("Kafka topic name is required.")

// ErrPartitionRequired happens when partition flag is not set.
var ErrPartitionRequired = errors.New("Kafka topic partition is required.")

// ErrOffsetRequired happens when offset flag is not set.
var ErrOffsetRequired = errors.New("Offset value is required.")

// SetBrokers holds the brokers flag value.
var SetBrokers string

// SetGroup holds the group flag value.
var SetGroup string

// SetTopic holds the topic flag value.
var SetTopic string

// SetPartition holds the partition flag value.
var SetPartition int

// SetOffset holds the offset flag value.
var SetOffset int

// SetAction is called after CLI flags are parsed and implements the actual "set offset" logic.
func SetAction(c *cli.Context) error {
	if SetBrokers == "" {
		return ErrBrokersRequired
	}

	if SetGroup == "" {
		return ErrGroupRequired
	}

	if SetTopic == "" {
		return ErrTopicRequired
	}

	if SetPartition == -1 {
		return ErrPartitionRequired
	}

	if SetOffset == -1 {
		return ErrOffsetRequired
	}

	config := kafka.NewConfig()
	config.BrokerList = strings.Split(SetBrokers, ",")
	client, err := kafka.New(config)
	if err != nil {
		return err
	}

	return client.CommitOffset(SetGroup, SetTopic, int32(SetPartition), int64(SetOffset))
}

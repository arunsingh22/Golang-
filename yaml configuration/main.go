package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Config map[string]interface{}

// type config struct {
// 	ProcessName      string `json:"process_name"`
// 	LogSeverity      string `json:"log_severity"`
// 	LogFilename      string `json:"log_filename"`
// 	VisionHost       string `json:"vision_host"`
// 	EventstoreConfig struct {
// 		DynamoEndpoint string `json:"dynamo_endpoint"`
// 		KafkaConfig    struct {
// 			Hosts               string   `json:"hosts"`
// 			CountProducers      string   `json:"count_producers"`
// 			CountRetries        string   `json:"count_retries"`
// 			CountPartitions     string   `json:"count_partitions"`
// 			Acks                string   `json:"acks"`
// 			PartitionerStrategy string   `json:"partitioner_strategy"`
// 			Topics              []string `json:"topics"`
// 		} `json:"kafka_config"`
// 	} `json:"eventstore_config"`
// }

func main() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("ERROR")
	}

	// fmt.Printf("%+v", Config)
	for k, v := range Config {
		fmt.Println(k, v)
	}

}

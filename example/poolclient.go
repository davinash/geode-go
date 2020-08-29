package main

import (
	"fmt"
	client "github.com/davinash/geode-go"
	"log"
)

func main() {
	geodeClient := &client.GeodeClient{}
	geodeClient.AddServer("127.0.0.1", 40404)

	region := geodeClient.Region("SampleData")
	log.Println("----- Put -----")
	for i := 0; i < 10; i++ {
		err := region.Put(
			fmt.Sprintf("Key-%d", i),
			fmt.Sprintf("Value-%d", i))
		if err != nil {
			log.Println(err)
		}
	}
}

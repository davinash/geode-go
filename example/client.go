package main

import (
	"fmt"
	client "github.com/davinash/geode-go"
	"log"
)

func main() {
	geodeClient, err := client.NewConnection("127.0.0.1", 40404)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Performing Put operation")
	for i := 0; i < 10; i++ {
		err = geodeClient.Region("SampleData").Put(
			fmt.Sprintf("Key-%d", i),
			fmt.Sprintf("Value-%d", i))
		if err != nil {
			log.Println(err)
		}
	}
}

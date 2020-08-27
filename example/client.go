package main

import (
	client "github.com/davinash/geode-go"
	"log"
)

func main() {
	geodeClient, err := client.NewConnection("127.0.0.1", 40404)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Performing Put operation")
	err = geodeClient.Region("SampleData").Put("K1", "V1")
	if err != nil {
		log.Println(err)
	}
}

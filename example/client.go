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
	region := geodeClient.Region("SampleData")
	log.Println("Performing Put operation")
	for i := 0; i < 10; i++ {
		err = region.Put(
			fmt.Sprintf("Key-%d", i),
			fmt.Sprintf("Value-%d", i))
		if err != nil {
			log.Println(err)
		}
	}
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("Key-%d", i)
		v, err := region.Get(k)
		if err == nil {
			log.Printf("Key = %v  Value = %v\n", k, v)
		}
	}
	val, err := region.PutIfAbsentRequest("key1", "value11")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("PutIfAbsentRequest -> %v", val)
}

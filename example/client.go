package main

import (
	"encoding/json"
	"fmt"
	client "github.com/davinash/geode-go"
	"github.com/davinash/geode-go/pkg"
	"log"
)

func main() {
	geodeClient, err := client.NewClient(100)
	if err != nil {
		log.Fatalln(err)
	}

	err = geodeClient.AddServer("127.0.0.1", 40404)
	if err != nil {
		log.Fatalln(err)
	}
	err = geodeClient.AddServer("127.0.0.1", 40405)
	if err != nil {
		log.Fatalln(err)
	}
	err = geodeClient.AddServer("127.0.0.1", 40406)
	if err != nil {
		log.Fatalln(err)
	}

	region := geodeClient.Region("SampleData")
	log.Println("----- Put -----")
	for i := 0; i < 10; i++ {
		err = region.Put(
			fmt.Sprintf("Key-%d", i),
			fmt.Sprintf("Value-%d", i))
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("----- Get -----")
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("Key-%d", i)
		v, err := region.Get(k)
		if err == nil {
			log.Printf("Key = %v  Value = %v\n", k, v)
		}
	}
	log.Println("----- PutIfAbsent -----")
	val, err := region.PutIfAbsent("key1", "value11")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("PutIfAbsent -> %v", val)

	log.Println("----- PutAll -----")
	kvs := make([]*pkg.KeyValue, 0)
	for i := 0; i < 10; i++ {
		kvs = append(kvs, &pkg.KeyValue{
			Key:   fmt.Sprintf("Key-PutAll-%d", i),
			Value: fmt.Sprintf("Key-PutAll-%d", i*100),
		})
	}
	_, err = region.PutAll(kvs)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("----- GetAll -----")
	keys := make([]string, 0)
	for i := 0; i < 10; i++ {
		keys = append(keys, fmt.Sprintf("Key-PutAll-%d", i))
	}
	//values, err := region.GetAll(keys)
	//for _, v := range values {
	//	log.Printf("Key = %v Value = %v\n", v.Key, v.Value)
	//}

	log.Println("----- Remove -----")
	err = region.Remove("Key-PutAll-0")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := region.Get("Key-PutAll-0")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Get Response after Remove = %v\n", resp)

	log.Println(geodeClient.GetRegionNames())
	size, err := region.Size()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Region size  %v\n", size)

	log.Println(region.KeySet())
	err = region.Clear()
	if err != nil {
		log.Fatalln(err)
	}
	ks, _ := region.KeySet()
	log.Printf("After Clearn Op = %v\n", ks)

	region.Put("abc", "xyz")
	_, _ = region.Get("abc")

	type Key struct {
		Name string
	}
	type Value struct {
		Age  int
		Dept string
	}
	err = region.Put(Key{
		Name: "Hello1",
	}, Value{
		Age:  10,
		Dept: "Dept1",
	})
	if err != nil {
		log.Println(err)
	}

	v, err := region.Get(Key{
		Name: "Hello1",
	})
	if err != nil {
		log.Printf("=-=-=-=-=-= : %v\n", err)
	}
	var v1 Value
	json.Unmarshal(v.([]byte), &v1)
	fmt.Println(v1)

}

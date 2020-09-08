[![Go Actions Status](https://github.com/davinash/geode-go/workflows/Go/badge.svg)](https://github.com/davinash/geode-go/actions)
# geode-go
geode-go provides Go package for connecting with [Apache Geode](https://geode.apache.org/).

#### Inspired by
* https://github.com/gemfire/geode-go-client

## Install
```bash
go get -u github.com/davinash/geode-go
```

## Region APIs Supported
* Put
* Get
* PutAll
* GetAll
* PutIfAbsent
* Remove
* Size
* Clear

## API Usage
#### Create a new client 
```go
geodeClient, err := client.NewClient(100)
```
#### Add a Server to this client
```go
err := geodeClient.AddServer("localhost", port)
if err != nil {
    return err
}
```
#### Put Operation
```go
geodeClient.Region(regionName).Put("Key", "Value")
```
#### Get Operation
```go
geodeClient.Region(regionName).Get("Key")
```
#### Size
```go
geodeClient.Region(regionName).Size()
```
#### KeySet
```go
geodeClient.Region(regionName).KeySet()
```

#### PutIfAbsent
```go
geodeClient.Region(regionName).PutIfAbsent("Key-1", "Value-New")
```
#### PutAll
```go
kvs := make([]*pkg.KeyValue, 0)
for i := 0; i < 10; i++ {
    kvs = append(kvs, &pkg.KeyValue{
        Key:   fmt.Sprintf("Key-PutAll-%d", i),
        Value: fmt.Sprintf("Val-PutAll-%d", i),
    })
}
putAll, err := suite.Client.Region(regionName).PutAll(kvs)
if err != nil {
    log.Println("PutAll failed, Error =  %v", err)
}
```
#### GetAll
```go
keys := make([]*pkg.Keys, 0)
for i := 0; i < 10; i++ {
    keys = append(keys, &pkg.Keys{
        Key: fmt.Sprintf("Key-PutAll-%d", i),
    })
}
values, err := suite.Client.Region(regionName).GetAll(keys)
if err != nil {
    log.Println("GetAll failed, Error = %v", err)
}
```
### Remove
```go
geodeClient.Region(regionName).Remove(key)
```
### Clear
```go
geodeClient.Region(regionName).Clear()
```
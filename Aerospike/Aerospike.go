package Aerospike

import (
	"fmt"

	"github.com/aerospike/aerospike-client-go"
)

//go get github.com/aerospike/aerospike-client-go

func getDataFromCache(key string) (string, error) {
	client, err := aerospike.NewClient("localhost", 3000)
	if err != nil {
		return "", err
	}

	defer client.Close()

	// Set up the key for the cache entry
	cacheKey, err := aerospike.NewKey("test", "cache", key)
	if err != nil {
		return "", err
	}

	// Retrieve the value from the cache
	rec, err := client.Get(nil, cacheKey)
	if err != nil {
		return "", err
	}

	// Extract the value from the record
	value, ok := rec.Bins["value"].(string)
	if !ok {
		return "", fmt.Errorf("value not found for key %s", key)
	}

	return value, nil
}

func storeDataInCache(key string, value string) error {
	client, err := aerospike.NewClient("localhost", 3000)
	if err != nil {
		return err
	}

	defer client.Close()

	// Set up the key and value for the cache entry
	cacheKey, err := aerospike.NewKey("test", "cache", key)
	if err != nil {
		return err
	}
	bins := aerospike.BinMap{
		"value": value,
	}

	// Store the cache entry in Aerospike
	err = client.Put(nil, cacheKey, bins)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Store data in the cache
	err := storeDataInCache("foo", "bar")
	if err != nil {
		panic(err)
	}

	// Retrieve data from the cache
	value, err := getDataFromCache("foo")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}

package Couchbase

import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

//go get gopkg.in/couchbase/gocb.v1

func getDataFromCache(key string) (string, error) {
	cluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		return "", err
	}

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "myusername",
		Password: "mypassword",
	})

	defer cluster.Close()

	bucket, err := cluster.OpenBucket("mybucket", "")
	if err != nil {
		return "", err
	}

	// Retrieve the value from the cache
	var value string
	_, err = bucket.Get(key, &value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func storeDataInCache(key string, value string) error {
	cluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		return err
	}

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "myusername",
		Password: "mypassword",
	})

	defer cluster.Close()

	bucket, err := cluster.OpenBucket("mybucket", "")
	if err != nil {
		return err
	}

	// Store the cache entry in Couchbase
	_, err = bucket.Upsert(key, value, 0)
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

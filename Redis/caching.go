package Redis

import (
    "github.com/go-redis/redis"
    "time"
)

type Cache struct {
    client *redis.Client
}

func NewCache() *Cache {
    return &Cache{
        client: redis.NewClient(&redis.Options{
            Addr:     "localhost:6379",
            Password: "", // no password set
            DB:       0,  // use default DB
        }),
    }
}

func (c *Cache) Set(key string, value string, expiration time.Duration) error {
    err := c.client.Set(key, value, expiration).Err()
    if err != nil {
        return err
    }
    return nil
}

func (c *Cache) Get(key string) (string, error) {
    val, err := c.client.Get(key).Result()
    if err == redis.Nil {
        return "", nil
    } else if err != nil {
        return "", err
    }
    return val, nil
}

func main() {
    cache := NewCache()

    err := cache.Set("foo", "bar", time.Minute)
    if err != nil {
        panic(err)
    }

    val, err := cache.Get("foo")
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}

/*
This example defines a Cache struct with Set and Get methods that use the Redis client to store 
and retrieve key-value pairs. The NewCache function creates a new Redis client with default options 
(connecting to a Redis server running on localhost:6379 with no password set).

In the Set method, the redis.Set command is used to store a key-value pair with an expiration time. 
In the Get method, the redis.Get command is used to retrieve the value associated with a given key. 
If the key is not found, the method returns an empty string and no error.

In the main function, a new cache is created using NewCache, and a value is stored using Set. 
The value is then retrieved using Get and printed to the console.
*/

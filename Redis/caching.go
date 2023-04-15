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

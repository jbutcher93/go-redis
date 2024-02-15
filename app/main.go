package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func RedisWriteClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		// Addr:     "redis-master.default.svc.cluster.local:6379", // stage and prod environments
		Addr:     "localhost:6380", // local environment
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0,
	})

	return client
}

func RedisReadClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		// Addr:     "redis-replicas.default.svc.cluster.local:6379",	// stage and prod environments
		Addr:     "localhost:6379", // local environment
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0,
	})

	return client
}

func incrementHandler(c *gin.Context) {
	client := RedisWriteClient()
	defer client.Close()

	counter, err := client.Incr(ctx, "visits").Result()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error incrementing visitor counter: %s", err)
		return
	}

	c.String(http.StatusOK, "Visitor number: %d", counter)
}

func get(c *gin.Context) {
	client := RedisReadClient()
	defer client.Close()

	key := c.Param("key")
	value := client.Get(ctx, key)

	if value.Err() != nil {
		c.String(http.StatusInternalServerError, "Error retrieving value from key: %s", key)
		return
	}

	c.String(http.StatusOK, "Value retrieved from key is: %s", value.Val())
}

func set(c *gin.Context) {
	var keyValue kv
	if err := c.BindJSON(&keyValue); err != nil {
		c.AbortWithStatus(400)
		return
	}

	client := RedisWriteClient()
	defer client.Close()

	err := client.Set(ctx, keyValue.Key, keyValue.Value, 0).Err()
	if err != nil {
		log.Print(errors.New("Failed to created new key-value pair"))
		c.AbortWithStatus(400)
		return
	}

	c.IndentedJSON(http.StatusCreated, keyValue)
}

func main() {
	router := gin.Default()
	router.GET("/", incrementHandler)
	router.GET("/get/:key", get)
	router.POST("/set", set)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}

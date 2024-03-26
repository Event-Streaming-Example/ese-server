package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ese.server/models"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client          *redis.Client
	context         context.Context
	expiryInMinutes int
}

func (r *RedisClient) Initialize(config Config) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Url, config.Port),
	})
	r.client = client
	r.context = context.Background()
	r.expiryInMinutes = config.ExpiryInMinutes
}

func (r *RedisClient) AddEvent(event models.Event) {
	jsonData, err := json.Marshal(event)
	if err != nil {
		log.Fatal("Error marshaling data:", err)
		return
	}

	// Use HSET to store JSON data in Redis hash
	err = r.client.HSet(r.context, event.Ip, fmt.Sprintf("%d", event.ClientTimestamp), jsonData).Err()
	if err != nil {
		log.Fatal("Error storing data in Redis:", err)
		return
	}

	// Set expiration for the key
	expirationTime := time.Duration(r.expiryInMinutes) * time.Minute
	err = r.client.Expire(r.context, event.Ip, expirationTime).Err()
	if err != nil {
		log.Fatal("Error setting expiration:", err)
		return
	}
}

func (r *RedisClient) AddEvents(events []models.Event) {
	// Start a pipeline
	pipe := r.client.Pipeline()

	// Iterate over events and add them to the pipeline
	for _, event := range events {
		jsonData, err := json.Marshal(event)
		if err != nil {
			log.Println("Error marshaling data:", err)
			continue
		}

		// Use HSET to store JSON data in Redis hash in the pipeline
		pipe.HSet(r.context, event.Ip, fmt.Sprintf("%d", event.ClientTimestamp), jsonData)

		// Set expiration for the key in the pipeline
		expirationTime := time.Duration(r.expiryInMinutes) * time.Minute
		pipe.Expire(r.context, event.Ip, expirationTime)
	}

	// Execute the pipeline
	_, err := pipe.Exec(r.context)
	if err != nil {
		log.Fatal("Error executing pipeline:", err)
		return
	}
}

func (r *RedisClient) EventExists(ip string) int64 {
	exists, err := r.client.Exists(r.context, ip).Result()
	if err != nil {
		log.Fatal("Error retrieving data from Redis:", err)
		return 0
	}
	return exists
}

func (r *RedisClient) GetAllEvents() []models.EventIpLog {
	var result = []models.EventIpLog{}
	var keys = r.getAllKeys()
	for _, key := range keys {
		var eventLogs = []models.Event{}
		// Use HGETALL to retrieve all fields and values from Redis hash
		fieldsValues, err := r.client.HGetAll(r.context, key).Result()
		if err != nil {
			log.Fatal("Error retrieving data from Redis:", err)
			return result
		}
		for _, jsonData := range fieldsValues {
			// Unmarshal JSON data back to Event
			var retrievedEvent models.Event
			err = json.Unmarshal([]byte(jsonData), &retrievedEvent)
			if err != nil {
				log.Fatal("Error unmarshaling data:", err)
				return result
			}
			eventLogs = append(eventLogs, retrievedEvent)
		}
		newEventIpLog := models.EventIpLog{
			Ip:        key,
			EventLogs: eventLogs,
		}
		result = append(result, newEventIpLog)
	}
	return result
}

func (r *RedisClient) getAllKeys() []string {
	var cursor uint64
	keys := make([]string, 0)

	for {
		var keysBatch []string
		var err error

		// Scan with a cursor to retrieve a batch of keys
		keysBatch, cursor, err = r.client.Scan(r.context, cursor, "*", 10).Result()
		if err != nil {
			log.Fatal("Error scanning keys:", err)
			return keys
		}

		// Append the batch of keys to the result
		keys = append(keys, keysBatch...)

		// Break the loop if the cursor is 0, indicating the end of the keys
		if cursor == 0 {
			break
		}
	}
	return keys
}

func (r *RedisClient) Destroy() {
	err := r.client.Close()
	if err != nil {
		log.Fatal("Error closing Redis client:", err)
		return
	}
	log.Println("Closed Redis Client")
}

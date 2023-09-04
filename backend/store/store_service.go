package store

import (
	"context"
	"fmt"

	"github.com/ckive/gourl/backend/constants"

	"github.com/redis/go-redis/v9"
)

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

type StorageService struct {
	redisClient *redis.Client
}

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// if user id was not provided generate one on the fly : case for not logged in users

/*
	We want to be able to save the mapping between the originalUrl

and the generated shortUrl url
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, 0).Err() // keep alive forever
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}

// Returns bool for if the customURL requested is in the cache
func CustomLinkInCache(shortUrl string) bool {
	exists, err := storeService.redisClient.Exists(ctx, shortUrl[:min(constants.ShortLinkLength, len(shortUrl))]).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed CustomLinkInCache url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	if exists == 1 {
		return true
	} else {
		return false
	}
}

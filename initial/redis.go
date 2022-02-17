package initial

import (
	"camp-backend/types"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const (
	redisAddr = "localhost:6379"
	password  = ""
)

var RedisContext = context.Background()
var RedisClient *redis.Client

func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       0,
	})

	pong, err := RedisClient.Ping(RedisContext).Result()
	fmt.Println("Redis ping result: ", pong)
	if err != nil {
		panic(fmt.Sprintf("redis ping failed, err is %s", err))
	}

	InsertDataToRedis()
}

func InsertDataToRedis() {
	students := make([]types.TMember, 0)
	Db.Select("user_id").Find(&students, "user_type = ? AND is_deleted = ?", 2, 0)
	for _, v := range students {
		RedisClient.SAdd(RedisContext, "students", v.UserID)
	}

	courses := make([]types.TCourse, 0)
	Db.Select("course_id", "capacity").Find(&courses)
	for _, v := range courses {
		RedisClient.SAdd(RedisContext, "courses", v.CourseID)
		RedisClient.HSet(RedisContext, "course:"+v.CourseID, "capacity", v.Capacity, "count", 0)
	}
}

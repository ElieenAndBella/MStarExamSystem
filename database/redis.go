/*
 * Created Date: Friday, July 8th 2022, 7:51:42 pm
 * Author: ZhanG
 */
package database

import (
	"fmt"

	"github.com/ElieenAndBella/MStarExamSystem/constant"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", constant.RC.Addr, constant.RC.Port),
		Password: constant.RC.Pass,
		DB:       constant.RC.DB,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}

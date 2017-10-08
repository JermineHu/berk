// This file "redis.go" is created by Lincan Li at 5/11/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package conf

import (
	"gopkg.in/redis.v3"
	"os"
)

func GetRedisOption() *redis.Options {
	addr := os.Getenv("REDIS_PORT_6379_TCP_ADDR") + ":" + os.Getenv("REDIS_PORT_6379_TCP_PORT")
	return &redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
}

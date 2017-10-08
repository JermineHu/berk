// This file "redis.go" is created by Lincan Li at 5/11/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package config

import (
	"gopkg.in/redis.v3"
)

func SetAddr(a string) {
	addr = a
}

var (
	rClient *redis.Client
	addr    string
)

func GetRedis() *redis.Client {
	if rClient == nil {
		rClient = connRedis()
	}
	return rClient
}

func connRedis() *redis.Client {
	rClient = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	return rClient
}

func SetRedis(ops *redis.Options) *redis.Client {
	rClient = redis.NewClient(ops)
	return rClient
}

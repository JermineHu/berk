// This file "token" is created by Lincan Li at 5/11/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package model

import (
	. "git.vdo.space/foy/config"
	"gopkg.in/redis.v3"
	"strconv"
)

func InsertUserToken(uID int64, token string) (string, error) {
	key := token
	val := strconv.Itoa(int(uID))
	err := GetRedis().Set(key, val, 0).Err()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return key, nil
}

func InsertUserTokenWithIDString(uID string, token string) (string, error) {
	err := GetRedis().Set(token, uID, 0).Err()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return token, nil
}

func GetUserIDByToken(token string) (int64, error) {
	key := token
	if token == "" {
		return 0, nil
	}
	val, err := GetRedis().Get(key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return int64(i), nil
}

func GetUserIDByTokenWithString(token string) (string, error) {
	val, err := GetRedis().Get(token).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}

func InsertDeviceToken(uID int64, dToken string) (string, error) {
	key := strconv.Itoa(int(uID))
	val := dToken

	err := GetRedis().Set(key, val, 0).Err()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return key, nil
}

func GetDeviceTokenByUser(uID int64) (string, error) {
	key := strconv.Itoa(int(uID))

	val, err := GetRedis().Get(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return val, nil
}

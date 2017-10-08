// This file "config.go" is created by Jermine Hu at 6/22/16.
// Copyright © 2016 - Jermine Hu. All rights reserved

package conf

import "os"

const (
	ZIPKIN_ADDRESS    = "ENV_ZIPKIN_ADDR"
	KAFKA_ADDR        = "ENV_KAFKA_ADDR"
	REDIS_ADDRESSES   = "ENV_REDIS_ADDR"
	REDIS_PORT        = "ENV_REDIS_PORT"
	MGO_ADDRESSES     = "ENV_MGO_ADDR"
	MGO_PORT          = "ENV_MGO_PORT"
	MGO_DATABASE      = "ENV_MGO_SMS_DATABASE"
	MGO_USERNAME      = "ENV_MGO_USERNAME"
	MGO_PASSWORD      = "ENV_MGO_PASSWORD"
	ENV_RESOURCE_ADDR = "ENV_RESOURCE_ADDR"
	ENV_RESOURCE_PATH = "ENV_RESOURCE_PATH"
)

func Conf_GetValue(key string) string {
	return conf.Get(key).String("")
}

func Conf_GetValues(key string) []string {
	var vals []string
	return conf.Get(key).StringSlice(vals)
}

func GetWXAppID() string {

	return os.Getenv("WX_AppID")

}

func GetWXAppSecret() string {

	return os.Getenv("WX_AppSecret")

}

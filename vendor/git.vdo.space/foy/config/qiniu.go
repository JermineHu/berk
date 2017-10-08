// This file "struct.go" is created by Lincan Li at 3/1/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package config

type QiNiuConfig struct {
	AccessKey      string
	SecretKey      string
	BucketName     string
	BucketURL      string
	CallbackDomain string
}

var (
	_QINIU_ACCESS_KEY      string
	_QINIU_SECRET_KEY      string
	_QINIU_BUCKET_NAME     string
	_QINIU_BUCKET_URL      string
	_QINIU_CALLBACK_DOMAIN string
)

func SetQINIU_ACCESS_KEY(s string) {
	_QINIU_ACCESS_KEY = s
}

func SetQINIU_SECRET_KEY(s string) {
	_QINIU_SECRET_KEY = s
}

func SetQINIU_BUCKET_NAME(s string) {
	_QINIU_BUCKET_NAME = s
}

func SetQINIU_BUCKET_URL(s string) {
	_QINIU_BUCKET_URL = s
}

func SetQINIU_CALLBACK_DOMAIN(s string) {
	_QINIU_CALLBACK_DOMAIN = s
}

var (
	q *QiNiuConfig
)

func GetQiNiuConfig() *QiNiuConfig {
	if q == nil {
		q = LoadQiNiuConfig()
	}
	return q
}

func LoadQiNiuConfig() *QiNiuConfig {

	return &QiNiuConfig{
		AccessKey:      _QINIU_ACCESS_KEY,
		SecretKey:      _QINIU_SECRET_KEY,
		BucketName:     _QINIU_BUCKET_NAME,
		BucketURL:      _QINIU_BUCKET_URL,
		CallbackDomain: _QINIU_CALLBACK_DOMAIN,
	}
}

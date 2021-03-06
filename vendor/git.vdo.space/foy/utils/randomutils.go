// This file "random.go" is created by Lincan Li at 1/25/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package utils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var numbers = []rune("0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func RandomNumberString(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = numbers[r.Intn(len(numbers))]
	}
	return string(b)
}

// 产生[start, end)的随机整数
func RandInt(intStart, intEnd int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := intStart + r.Intn(intEnd-intStart)
	return ret
}

// 产生(ageStart, ageEnd)的生日
func RandBirth(ageStart, ageEnd int) time.Time {
	year, _, _ := time.Now().Date()
	y := year - RandInt(ageStart, ageEnd)
	m := RandInt(1, 13)

	var d int
	if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		d = RandInt(1, 32)
	} else if m == 2 {
		d = RandInt(1, 29)
	} else {
		d = RandInt(1, 31)
	}
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	return t
}

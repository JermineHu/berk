package utils

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"sort"
	"strconv"
	"time"
)

func Str2UUID(a string) uuid.UUID {
	return uuid.FromStringOrNil(a)
}

func Strs2UUIDs(as []string) []uuid.UUID {
	us := make([]uuid.UUID, len(as))

	for k, v := range as {
		us[k] = uuid.FromStringOrNil(v)
	}
	return us
}

func UUIDs2Strs(us []uuid.UUID) []string {
	ss := make([]string, len(us))

	for k, v := range us {
		ss[k] = v.String()
	}
	return ss
}

func Str2Time(ts string) time.Time {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return time.Time{}
	}
	return t
}

func Time2Str(t time.Time) string {
	if t == (time.Time{}) || t.IsZero() {
		return ""
	}

	ts := t.Format(time.RFC3339)
	return ts
}

func StoI(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	a, err := strconv.Atoi(s)
	return a, err
}

func StoB(s string) (bool, error) {
	if s == "true" {
		return true, nil
	}
	return false, nil
}

func StoI64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	a, err := strconv.ParseInt(s, 10, 64)
	return a, err
}

func ItoStr(i int) string {
	a := strconv.Itoa(i)
	return a
}

func Str2ObjectId(id string) bson.ObjectId {

	var oid bson.ObjectId

	oid = bson.ObjectIdHex(id)

	return oid
}

func Strs2ObjectIds(ids []string) []*bson.ObjectId {

	oids := []*bson.ObjectId{}

	for _, i := range ids {

		oid := bson.ObjectIdHex(i)

		oids = append(oids, &oid)
	}

	return oids
}

func ObjectIds2Strs(ids []*bson.ObjectId) []string {

	strids := []string{}

	for _, i := range ids {

		strids = append(strids, i.Hex())
	}

	return strids
}

// RemoveDuplicatesAndEmpty 去除重复和空字符串
func RemoveDuplicatesAndEmpty(s []string) (ret []string) {
	sort.Strings(s)
	a_len := len(s)
	for i := 0; i < a_len; i++ {
		if (i > 0 && s[i-1] == s[i]) || len(s[i]) == 0 {
			continue
		}
		ret = append(ret, s[i])
	}
	return
}

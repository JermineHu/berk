package model

import (

	//"gopkg.in/mgo.v2"
	//"strconv"
	//"time"
	//"git.vdo.space/foy/config"
)



//func GeneratorNumForLottery(MDB *mgo.Database) string {
//
//	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	var nums = []rune("0123456789")
//
//	currnum, _ := config.GetAutoIncID("curr_num", MDB)
//	currlet, _ := config.GetAutoID("curr_letter", MDB)
//
//	currPicNum, _ := config.GetAutoIncID("curr_pic_num", MDB)
//
//	if currnum.ID == int64(len(nums)) {
//		currlet, _ = config.GetAutoIncID("curr_letter", MDB)
//		currnum, _ = config.UpdataAutoIncID("curr_num", MDB, 0)
//
//	}
//
//	if currlet.ID == int64(len(letters)) {
//		currlet, _ = config.UpdataAutoIncID("curr_letter", MDB, 0)
//	}
//
//	now := time.Now()
//
//	rstr := strconv.FormatInt(currPicNum.ID, 10)
//
//	b := make([]rune, 2)
//
//	b[0] = letters[currlet.ID]
//	b[1] = nums[currnum.ID]
//
//	d := strconv.FormatInt(int64(now.Day()), 10)
//	h := strconv.FormatInt(int64(now.Hour()), 10)
//
//	return string(b) + d + h + rstr
//
//}

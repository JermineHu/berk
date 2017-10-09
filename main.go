package main

import (
	"git.vdo.space/berk/conf"
	ml "git.vdo.space/berk/model"
	"git.vdo.space/berk/route"
	dc "git.vdo.space/foy/config"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var MDB *mgo.Database

func init() {
	//os.Setenv("MONGODB_CONNECTION", "Jermine:123456@localhost:27018/berk")
	//os.Setenv("REDIS_PORT", "127.0.0.1:6379")
	MDB = conf.GetMongo()
	ml.InitMDB(MDB)
	dc.SetRedis(conf.GetRedisOption())
}

func main() {
	//now := time.Now()
	//
	//dt, _ := time.Parse("2006-01-02", now.Format("2006-01-02"))
	//durdm, _ := time.ParseDuration("7h")
	//tm := dt.Add(durdm).Unix() - now.Unix() //计算距离下次启动还有多久

	ticker := time.NewTicker(time.Duration(8) * time.Hour)
	go func() {
		for _ = range ticker.C {

			client := &http.Client{}
			req, err := http.NewRequest("POST", "https://openapi.daocloud.io/v1/apps/"+os.Getenv("AppID")+"/actions/restart", strings.NewReader(""))
			if err != nil {
				// handle error
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", os.Getenv("AppAuthorization"))
			resp, err := client.Do(req)
			defer resp.Body.Close()

		}
	}()

	mEngine := route.Mean{}.Engine()
	//log.Fatal("HTTP start fail ", mEngine.RunTLS(":8080","/certs/server.crt","/certs/server.key"))
	log.Fatal("HTTP start fail ", mEngine.Run(":8000"))

}

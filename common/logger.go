package common

import (
	log "github.com/Sirupsen/logrus"
	//	"github.com/weekface/mgorus"
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const LOG_CONFIG = "logconfig.ini" //用于存放当前正在写入的日志文件名
var lastLogFileName string         //最后一个日志文件名称

type MeanLogger struct {
	RequestID uuid.UUID
}

type LogEntry struct {
	*log.Entry
}

func NewMeanLogger() *MeanLogger {
	return &MeanLogger{
		RequestID: uuid.NewV4(),
	}
}

var sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)

func (l *MeanLogger) Print(values ...interface{}) {
	level := values[0]
	currentTime := "\n\033[33m[" + time.Now().Format("2006-01-02 15:04:05") + "]\033[0m"
	source := fmt.Sprintf("\033[35m(%v)\033[0m", values[1])
	messages := []interface{}{source, currentTime}

	if level == "sql" {
		// duration
		messages = append(messages, fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
		// sql
		var formatedValues []interface{}
		for _, value := range values[4].([]interface{}) {
			indirectValue := reflect.Indirect(reflect.ValueOf(value))
			if indirectValue.IsValid() {
				value = indirectValue.Interface()
				if t, ok := value.(time.Time); ok {
					formatedValues = append(formatedValues, fmt.Sprintf("'%v'", t.Format(time.RFC3339)))
				} else if b, ok := value.([]byte); ok {
					formatedValues = append(formatedValues, fmt.Sprintf("'%v'", string(b)))
				} else if r, ok := value.(driver.Valuer); ok {
					if value, err := r.Value(); err == nil && value != nil {
						formatedValues = append(formatedValues, fmt.Sprintf("'%v'", value))
					} else {
						formatedValues = append(formatedValues, "NULL")
					}
				} else {
					formatedValues = append(formatedValues, fmt.Sprintf("'%v'", value))
				}
			} else {
				formatedValues = append(formatedValues, fmt.Sprintf("'%v'", value))
			}
		}
		messages = append(messages, fmt.Sprintf(sqlRegexp.ReplaceAllString(values[3].(string), "%v"), formatedValues...))
	} else {
		messages = append(messages, "\033[31;1m")
		messages = append(messages, values[2:]...)
		messages = append(messages, "\033[0m")
	}

	logType := values[0]

	entry := l.WithFields(log.Fields{
		"module": "gorm",
		"type":   logType,
	})

	entry.Println(messages)
}

// Instance a Logger middleware with the specified writter buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func LoggerWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		//statusColor := colorForStatus(statusCode)
		//methodColor := colorForMethod(method)
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		//jsonRequest,_:=json.Marshal(c.Request)

		//fmt.Println(c.Accepted,c.Errors,c.Err(),c.Params,jsonRequest)

		//gin的请求日志:
		log.SetFormatter(&log.TextFormatter{})

		// A common pattern is to re-use fields between logging statements by re-using
		// the logrus.Entry returned from WithFields()
		contextLogger := log.WithFields(log.Fields{
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     method,
			"statusCode": statusCode,
			"comment":    comment,
			"path":       path,
		})

		contextLogger.Level = log.FatalLevel
		contextLogger.Message = "This log is about http request of gin!"
		contextLogger.Time = time.Now()

		//	byt, _ := contextLogger.Reader()

		//	ToEchoLogFile("logs", byt.Bytes(), 1024 * 1024 * 5)

	}

}

func (l *MeanLogger) Println(v ...interface{}) {

}

func (l *MeanLogger) WithFields(fields log.Fields) *LogEntry {
	fields["request-id"] = l.RequestID
	entry := log.WithFields(fields)
	return &LogEntry{entry}
}

func NewLogger() *MeanLogger {
	//	hooker, err := mgorus.NewHooker("localhost:27017", "db", "collection")
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	log.AddHook(hooker)

	return NewMeanLogger()
}

//getLastLogFileName 获取当前日志文件名称
func getLastLogFileName(path string) string {

	if len(lastLogFileName) != 0 {

		return lastLogFileName
	}

	logcfgName := path + "/" + LOG_CONFIG

	_, err := os.Stat(logcfgName)
	if os.IsNotExist(err) {
		setLastLogFileName(path, getNewLogFileName())
		return lastLogFileName
	} else {

		str, err := ioutil.ReadFile(logcfgName)
		if err != nil {
			panic(err)
		}
		lastLogFileName = string(str)
		return lastLogFileName
	}

}

//setLastLogFileName  设置当前日志文件名称
func setLastLogFileName(path string, value string) {

	logcfgName := path + "/" + LOG_CONFIG

	err := os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(logcfgName, os.O_RDWR|os.O_CREATE, 0644) //打开文件
	if err != nil {
		panic(err)
	}
	if _, err = f.WriteString(value); err != nil {
		panic(err)
	}

	lastLogFileName = value

}

//getNewLogFileName 创建新日志的名称
func getNewLogFileName() string {

	now := time.Now()
	var m time.Month = now.Month()
	var i int = int(m)

	fileName := strconv.Itoa(now.Year()) + "-" + strconv.Itoa(i) + "-" + strconv.Itoa(now.Day()) + " " + strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute()) + ".log"

	return fileName
}

//getCurrentDirectory 获取当前目录
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//ToEchoLogFile 将日志打印到文件
func ToEchoLogFile(path string, log []byte, fileSizeLimit int64) {

	logFileName := getLastLogFileName(path)
	fileInfo, err := os.Stat(logFileName)

	if !os.IsNotExist(err) || (fileInfo != nil && fileInfo.Size() >= fileSizeLimit) {

		logFileName = getNewLogFileName()
		setLastLogFileName(path, logFileName)

	}

	f, err := os.OpenFile(path+"/"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) //打开文件
	if err != nil {
		panic(err)
	}

	_, err = f.Write(log)

	if err != nil {
		panic(err)
	}

}

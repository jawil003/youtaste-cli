package logger

import (
	"bs-to-scrapper/server/enums"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type Instance struct {
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
}

func Logger() *Instance {

	f, err := os.OpenFile(enums.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	writer := io.MultiWriter(f, os.Stdout)

	config := log.LstdFlags | log.Lshortfile

	return &Instance{
		Info:  log.New(writer, "INFO: ", config),
		Error: log.New(writer, "ERROR: ", config),
		Warn:  log.New(writer, "WARN: ", config),
	}
}

func ConvertToString(inst interface{}) string {

	var result string

	v := reflect.ValueOf(inst)

	for i := 0; i < v.NumField(); i++ {

		varName := v.Type().Field(i).Name
		varValue := v.Field(i).Interface()

		result += fmt.Sprintf("%v=%v ", varName, varValue)
	}

	return result

}

func LogResponse(status int, response interface{}) string {

	res := fmt.Sprintf("Response is %d and body is weekdays=%v", status, ConvertToString(response))

	return res
}

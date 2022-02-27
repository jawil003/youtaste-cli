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

	t := reflect.TypeOf(inst)

	v := reflect.ValueOf(inst)

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {

			if i == 0 {
				result += "\n{\n"
			}

			f := t.Field(i)

			if f.Type.Kind() == reflect.Func {
				continue
			}

			result += fmt.Sprintf("   %s: %v\n", f.Name, v.Field(i))

			if i == t.NumField()-1 {
				result += "}\n"
			}

		}
	}

	return result

}

func LogResponse(status int, response interface{}) string {

	res := fmt.Sprintf("Response is %d and body is weekdays=%v", status, ConvertToString(response))

	return res
}

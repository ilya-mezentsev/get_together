package logger

import (
	"encoding/json"
	"fmt"
	"log"
)

type Fields struct {
	MessageTemplate string
	Args            []interface{}
	Optional        map[string]interface{}
}

func WithFields(fields Fields, logFn func(interface{})) {
	optionalArgs, _ := json.Marshal(fields.Optional)
	message := fmt.Sprintf(fields.MessageTemplate, fields.Args...)

	logFn(fmt.Sprintf("%s\n\t%s", message, "[CONTEXT]: "+string(optionalArgs)))
}

func Info(msg interface{}) {
	log.Println(
		"[INFO]: " + fmt.Sprintf("%v", msg))
}

func Warning(msg interface{}) {
	log.Println(
		"[WARNING]: " + fmt.Sprintf("%v", msg))
}

func WarningF(template string, a ...interface{}) {
	Warning(fmt.Sprintf(template, a...))
}

func Error(msg interface{}) {
	log.Println(
		"[ERROR]: " + fmt.Sprintf("%v", msg))
}

func ErrorF(template string, a ...interface{}) {
	Error(fmt.Sprintf(template, a...))
}

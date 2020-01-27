package logger

import (
  "fmt"
  "log"
)

var buffer string

func getMsgWithBuffer(msg string) string {
  if buffer == "" {
    return msg
  } else {
    res := buffer + " " + msg
    buffer = ""
    return res
  }
}

func Info(msg interface{}) {
  log.Println(
    "[INFO]: " + getMsgWithBuffer(fmt.Sprintf("%v", msg)))
}

func Warning(msg interface{}) {
  log.Println(
    "[WARNING]: " + getMsgWithBuffer(fmt.Sprintf("%v", msg)))
}

func WarningF(template string, a ...interface{}) {
  Warning(fmt.Sprintf(template, a...))
}

func Error(msg interface{}) {
  log.Println(
    "[ERROR]: " + getMsgWithBuffer(fmt.Sprintf("%v", msg)))
}

func ErrorF(template string, a ...interface{}) {
  Error(fmt.Sprintf(template, a...))
}

package clog

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"strconv"
)

var ClogItem *Clog

type Clog struct{}

func ClogGet() *Clog {
	if ClogItem == nil {
		ClogItem = &Clog{}
		return ClogItem
	} else {
		return ClogItem
	}
}

func Println(str interface{}) {
	msg := des(str)
	msg = fmt.Sprintf("%c[1;32;40m[%v]%c[0m %v", 0x1B, " INFO ", 0x1B, msg)

	log.Println(msg)
}

func PrintEr(str interface{}) {
	msg := des(str)
	msg = fmt.Sprintf("%c[1;33;40m[%v]%c[0m %v", 0x1B, " ERROR ", 0x1B, msg)

	log.Println(msg)
}

func PrintWa(str interface{}) {
	msg := des(str)
	msg = fmt.Sprintf("%c[1;31;40m[%v]%c[0m %v", 0x1B, " WARNING ", 0x1B, msg)

	log.Println(msg)
}

func Sprint(str string) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "[" + filename + ":" + strconv.Itoa(line) + "] " + str
	return msg
}

func des(str interface{}) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)

	msg := fmt.Sprintf("[ %v : %v ]  %v", filename, strconv.Itoa(line), str)
	return msg
}

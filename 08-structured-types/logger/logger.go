package logger

import (
	"log"
	"runtime"
)

func LogFatal(err error) {
	pc, _, line, _ := runtime.Caller(1)

	log.Fatalf("[ERROR] %s (func: %s, line: %d)", err, runtime.FuncForPC(pc).Name(), line)
}

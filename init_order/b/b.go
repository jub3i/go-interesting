package b

import (
	"log"
	"runtime"
)

func init() {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		log.Printf("[INIT] from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
}

func Fn() {
	log.Println("b.go")
}

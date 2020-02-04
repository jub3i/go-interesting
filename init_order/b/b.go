package b

import (
	"log"
	"runtime"
)

// first init function
func init() {
	log.Println("b.go - init_0")
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		log.Printf("[INIT] from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
}

// you can have multiple init functions
func init() {
	log.Println("b.go - init_1")
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		log.Printf("[INIT] from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
}

func Fn() {
	log.Println("b.go")
}

package main

// imports exec first, and their init fn gets called first
import (
	"go-interesting/init_order/b"
	"go-interesting/init_order/c"
	"log"
	"runtime"
)

// main.go's init function is called last
func init() {
	log.Println("main.go - init_0")
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		log.Printf("[INIT] from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}
}

func main() {
	log.Println("main.go")
	b.Fn()
	c.Fn()
}

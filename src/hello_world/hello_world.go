package main

//#include"foo.h"
import "C"
import "log"

func main() {
	var a = C.int(1)
	var b = C.int(2)
	var c = C.add(a, b)
	log.Println("result:", c)
}

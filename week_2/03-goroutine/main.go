package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello ")
}

func main() {
	go sayHello()           // yangi goroutine ishga tushdi
	time.Sleep(time.Second) // main tugab ketmasligi uchun kutamiz
}

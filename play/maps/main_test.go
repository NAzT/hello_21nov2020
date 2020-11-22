package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	total := 10
	now := time.Now()
	for i := 0; i < total; i++ {
		go write("A")
	}
	for i := 0; i < total; i++ {
		go read()
	}
	fmt.Println(time.Now().Sub(now))

}

var s string
var mux sync.Mutex

func main() {
}

func write(str string) {
	mux.Lock()
	s = str
	mux.Unlock()
}

func read() string {
	mux.Lock()
	defer mux.Unlock()
	return s
}

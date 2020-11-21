package main

import (
	"fmt"
	"strconv"
)

type Integer int

func (i Integer) toString() string {
	return strconv.Itoa(int(i))
}

func (i *Integer) set(n int) {
	*i = Integer(n)
}
func (i Integer) get() int {
	return int(i)
}

func main() {
	var i Integer = 5
	fmt.Printf("%q\n", i.toString())

	i.set(10)
	n := i.get()

	fmt.Println(n)
}

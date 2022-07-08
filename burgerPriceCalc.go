package main

import (
	"fmt"
)

func main() {
	a := burger{100}
	fmt.Println(a)
}

type burger struct {
	price int
}
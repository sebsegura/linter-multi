package main

import (
	"fmt"
	"sebsegura/foo/pkg/adder"
)

const ID = "example"

func main() {
	fmt.Println(ID)
	fmt.Println(adder.Add(2, 2))
}

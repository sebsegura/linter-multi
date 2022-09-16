package main

import (
	"fmt"
	"sebsegura/foo/pkg/adder"
)

const Id = "example"

func main() {
	fmt.Println(Id)
	fmt.Println(adder.Add(2, 2))
}

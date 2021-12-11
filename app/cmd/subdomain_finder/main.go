package main

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
)

func main() {
	fmt.Println("Start...")
	r := tarantool.DomainRepo{}
	fmt.Println(r)
}

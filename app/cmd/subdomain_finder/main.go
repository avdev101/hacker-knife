package main

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/adapters"
	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	fmt.Println("Start...")
	r := tarantool.SubdomainRepo{}
	q := tarantool.Queue{}
	f := adapters.DummySubdomainFinder{}

	s := core.DomainEnumerateService{&r, &f, &q}

	err := s.Enumerate("hackerone.com", false)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

package main

import (
	"fmt"
	"log"

	"github.com/eremeevdev/hacker-knife/adapters"
	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	q, err := tarantool.NewQueue("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	r, err := tarantool.NewSubdomainRepo("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	f := adapters.DummySubdomainFinder{}

	s := core.DomainEnumerateService{&r, &f, &q}

	for {
		task, err := q.TakeSubdomain()
		if err != nil {
			panic(err)
		}

		fmt.Println(task)

		err = s.Enumerate(task.Data.DomainName, task.Data.Propagate)
		if err != nil {
			log.Fatal(err)
			task.Nack()
		} else {
			task.Ack()
		}
	}

}

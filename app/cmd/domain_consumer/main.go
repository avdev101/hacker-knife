package main

import (
	"fmt"
	"log"

	"github.com/eremeevdev/hacker-knife/adapters/collectors/domain"
	"github.com/eremeevdev/hacker-knife/adapters/queue"
	"github.com/eremeevdev/hacker-knife/adapters/repo"

	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	q, err := queue.NewQueue("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	r, err := repo.NewSubdomainRepo("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	f := domain.DummySubdomainCollector{}

	s := core.DomainCollectService{&r, &f, &q}

	for {
		task, err := q.TakeCollectSubdomain()
		if err != nil {
			panic(err)
		}

		fmt.Println(task)

		err = s.Collect(task.Data.DomainName, task.Data.Propagate)
		if err != nil {
			log.Fatal(err)
			task.Nack()
		} else {
			task.Ack()
		}
	}

}

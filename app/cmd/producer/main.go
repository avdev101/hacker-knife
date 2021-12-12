package main

import (
	"github.com/eremeevdev/hacker-knife/adapters/queue"
	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	q, err := queue.NewQueue("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	t := core.CollectSubDomainCommand{DomainName: "hackerone.com"}

	err = q.CollectSubdomain(t)
	if err != nil {
		panic(err)
	}
}

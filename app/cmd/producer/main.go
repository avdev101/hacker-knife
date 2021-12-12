package main

import (
	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	q, err := tarantool.NewQueue("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	t := core.FidnSubDomainTask{DomainName: "hackerone.com"}

	err = q.FindSubdomain(t)
	if err != nil {
		panic(err)
	}
}

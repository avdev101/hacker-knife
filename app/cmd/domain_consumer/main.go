package main

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
)

func main() {
	q, err := tarantool.NewQueue("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	task, err := q.TakeSubdomain()
	if err != nil {
		panic(err)
	}

	fmt.Println(task)
	task.Ack()
}

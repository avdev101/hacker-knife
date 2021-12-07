package main

import (
	"fmt"

	"github.com/tarantool/go-tarantool"
	"github.com/tarantool/go-tarantool/queue"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Domain struct {
	Name   string
	Parent string
}

func (domain *Domain) DecodeMsgpack(d *msgpack.Decoder) error {
	m, err := d.DecodeMap()

	fmt.Println(m)
	x := m.(map[interface{}]interface{})["Name"]

	domain.Name = x.(string)

	return err

}

func main() {

	opts := tarantool.Opts{User: "admin", Pass: "pass"}
	conn, err := tarantool.Connect("tarantool_app:3722", opts)

	if err != nil {
		panic(err)
	}

	resp, err := conn.Select("domain", "primary", 0, 1, tarantool.IterEq, []interface{}{"ya.ru"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Code)
	fmt.Println(resp.Data)

	que := queue.New(conn, "parse_subdomain")

	//d := Domain{"hello.ru"}
	/*task, err := que.Put(d)
	if err != nil {
		panic(err)
	}

	fmt.Println(task.Data())*/
	//task, err := que.Take()

	fmt.Println("===========")

	var d Domain

	task, err := que.TakeTyped(&d)
	if err != nil {
		panic(err)
	}

	err = task.Release()
	if err != nil {
		panic(err)
	}

	fmt.Println(task.Data())
	fmt.Println(d)

}

package main

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/adapters/tarantool"
	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	fmt.Println("Start...")

	r, err := tarantool.NewSubdomainRepo("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	xDomains := []core.Subdomain{
		{Domain: "hackerone.com", Name: "api.hackerone.com"},
	}

	err = r.CreateBatch(xDomains)
	if err != nil {
		panic(err)
	}

	items, err := r.GetList("hackerone.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(items)

	/*domains := []core.Subdomain{
		{Name: "api.hackerone.com"},
	}

	err = r.DeleteBatch(domains)
	if err != nil {
		panic(err)
	}*/

	/*q := tarantool.Queue{}
	f := adapters.DummySubdomainFinder{}

	s := core.DomainEnumerateService{&r, &f, &q}

	err = s.Enumerate("hackerone.com", false)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)*/
}

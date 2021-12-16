package main

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/adapters/collectors/domain"
	"github.com/eremeevdev/hacker-knife/adapters/queue"
	"github.com/eremeevdev/hacker-knife/adapters/repo"

	"github.com/eremeevdev/hacker-knife/core"
)

func main() {
	fmt.Println("Start...")

	r, err := repo.NewSubdomainRepo("tarantool_app:3722", "admin", "pass")
	if err != nil {
		panic(err)
	}

	/*xDomains := []core.Subdomain{
		{ParentName: "hackerone.com", Name: "api.hackerone.com", IsNew: true},
	}

	err = r.UpdateBatch(xDomains)
	if err != nil {
		panic(err)
	}

	items, err := r.GetList("hackerone.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
	*/

	/*domains := []core.Subdomain{
		{Name: "api.hackerone.com"},
	}

	err = r.DeleteBatch(domains)
	if err != nil {
		panic(err)
	}*/

	q := queue.Queue{}
	f := domain.DummySubdomainCollector{}

	s := core.DomainCollectService{&r, &f, &q}

	err = s.Collect("hackerone.com", false)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

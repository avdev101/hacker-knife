package tarantool

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/core"
	"github.com/tarantool/go-tarantool"
)

type SubdomainRepo struct {
	conn *tarantool.Connection
}

func NewSubdomainRepo(host string, user string, pass string) (SubdomainRepo, error) {
	conn, err := tarantool.Connect(host, tarantool.Opts{
		User: user,
		Pass: pass,
	})

	if err != nil {
		return SubdomainRepo{}, fmt.Errorf("can't create connection: %v", err)
	}

	return SubdomainRepo{conn}, nil
}

func tuplesToSubdomains(tuples [][]interface{}) []core.Subdomain {
	result := make([]core.Subdomain, 0)

	for _, tuple := range tuples {
		item := core.Subdomain{
			Domain: tuple[0].(string),
			Name:   tuple[1].(string),
			Cname:  tuple[2].(string),
			IsNew:  tuple[3].(bool),
		}
		result = append(result, item)
	}

	return result
}

func (r *SubdomainRepo) GetList(domainName string) ([]core.Subdomain, error) {

	resp, err := r.conn.Select("subdomain", "domain", 0, 10000, tarantool.IterEq, []interface{}{domainName})
	if err != nil {
		return nil, err
	}

	tuples := resp.Tuples()

	subdomains := tuplesToSubdomains(tuples)

	return subdomains, nil
}

func (r *SubdomainRepo) DeleteBatch(domains []core.Subdomain) error {
	return nil
}

func (r *SubdomainRepo) CreateBatch(domains []core.Subdomain) error {
	return nil
}

func (r *SubdomainRepo) UpdateBatch(domains []core.Subdomain) error {
	return nil
}

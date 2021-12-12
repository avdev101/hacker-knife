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
	names := make([]string, len(domains))

	for i, d := range domains {
		names[i] = d.Name
	}

	_, err := r.conn.Call("batch_subdomain_delete", []interface{}{names})

	return err

}

func (r *SubdomainRepo) CreateBatch(domains []core.Subdomain) error {
	tuples := subdomainToTuples(domains)

	_, err := r.conn.Call("batch_subdomain_create", []interface{}{tuples})

	return err

}

func (r *SubdomainRepo) UpdateBatch(domains []core.Subdomain) error {
	return nil
}

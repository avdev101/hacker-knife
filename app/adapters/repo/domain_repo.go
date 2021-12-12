package repo

import (
	"github.com/eremeevdev/hacker-knife/core"
	"github.com/tarantool/go-tarantool"
)

type DomainRepo struct {
	conn *tarantool.Connection
}

func (r *DomainRepo) Create(d core.Domain) error {
	_, err := r.conn.Insert("domain", domainToTuple(d))
	return err
}

func (r *DomainRepo) GetList() ([]core.Domain, error) {

	resp, err := r.conn.Select("domain", "primary", 0, 10000, tarantool.IterGe, []interface{}{})

	if err != nil {
		return nil, err
	}

	tuples := resp.Tuples()

	domains := tuplesToDomains(tuples)

	return domains, nil

}

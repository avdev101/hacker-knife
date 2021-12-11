package tarantool

import "github.com/eremeevdev/hacker-knife/core"

type DomainRepo struct {
}

func (r *DomainRepo) Create(d core.Domain) error {
	return nil
}

func (r *DomainRepo) GetList() ([]core.Domain, error) {
	result := make([]core.Domain, 0)

	return result, nil
}

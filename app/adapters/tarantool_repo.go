package adapters

import "github.com/eremeevdev/hacker-knife/core"

type TarantoolDomainRepo struct {
}

func (r *TarantoolDomainRepo) Create(d core.Domain) error {
	return nil
}

func (r *TarantoolDomainRepo) GetList() ([]core.Domain, error) {
	result := make([]core.Domain, 0)

	return result, nil
}

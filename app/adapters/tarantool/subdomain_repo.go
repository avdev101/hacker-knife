package tarantool

import "github.com/eremeevdev/hacker-knife/core"

type SubdomainRepo struct {
}

func (r *SubdomainRepo) GetList(domainName string) ([]core.Subdomain, error) {
	result := make([]core.Subdomain, 0)
	return result, nil
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

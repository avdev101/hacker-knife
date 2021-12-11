package adapters

import "github.com/eremeevdev/hacker-knife/core"

type DummySubdomainFinder struct {
}

func (f *DummySubdomainFinder) Enumerate(domain string) ([]core.SubdomainFindItem, error) {
	result := make([]core.SubdomainFindItem, 0)
	result = append(result, core.SubdomainFindItem{"ya.ru", "x.ya.ru", "xx"})
	result = append(result, core.SubdomainFindItem{"ya.ru", "y.ya.ru", "yy"})

	return result, nil
}

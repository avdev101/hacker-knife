package adapters

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/core"
)

type DummySubdomainFinder struct {
}

func (f *DummySubdomainFinder) Enumerate(domain string) ([]core.SubdomainFindItem, error) {
	result := make([]core.SubdomainFindItem, 0)
	result = append(result, core.SubdomainFindItem{domain, fmt.Sprintf("%v.%v", "x", domain), "xx"})
	result = append(result, core.SubdomainFindItem{domain, fmt.Sprintf("%v.%v", "y", domain), "yy"})

	return result, nil
}

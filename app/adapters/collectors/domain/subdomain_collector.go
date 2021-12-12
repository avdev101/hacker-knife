package domain

import (
	"fmt"

	"github.com/eremeevdev/hacker-knife/core"
)

type DummySubdomainCollector struct {
}

func (f *DummySubdomainCollector) Collect(domain string) ([]core.SubdomainCollectItem, error) {
	result := make([]core.SubdomainCollectItem, 0)
	result = append(result, core.SubdomainCollectItem{domain, fmt.Sprintf("%v.%v", "x", domain), "xx1"})
	result = append(result, core.SubdomainCollectItem{domain, fmt.Sprintf("%v.%v", "y", domain), "yy1"})

	return result, nil
}

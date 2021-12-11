package core

type domainDiff struct {
	existing    []Subdomain
	existingMap map[string]Subdomain

	found    []SubdomainFindItem
	foundMap map[string]SubdomainFindItem
}

func getEMap(existing []Subdomain) map[string]Subdomain {
	result := make(map[string]Subdomain)

	for _, d := range existing {
		result[d.Name] = d
	}

	return result
}

func getFMap(found []SubdomainFindItem) map[string]SubdomainFindItem {
	result := make(map[string]SubdomainFindItem)

	for _, d := range found {
		result[d.Name] = d
	}

	return result
}

func newDomainDiff(existing []Subdomain, found []SubdomainFindItem) domainDiff {
	eMap := getEMap(existing)
	fMap := getFMap(found)

	return domainDiff{existing, eMap, found, fMap}

}

func (d *domainDiff) getNew() []Subdomain {
	return nil
}

func (d *domainDiff) getChanged() []Subdomain {
	return nil
}

func (d *domainDiff) getDeleted() []Subdomain {
	return nil
}

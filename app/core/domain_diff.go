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
	return domainDiff{existing, getEMap(existing), found, getFMap(found)}

}

func (d *domainDiff) getNew() []Subdomain {
	result := make([]Subdomain, 0)

	for _, found := range d.found {

		_, ok := d.existingMap[found.Name]

		if !ok {
			new := Subdomain{
				Domain: found.Domain,
				Name:   found.Name,
				Cname:  found.CName,
				IsNew:  true,
			}
			result = append(result, new)
		}

	}

	return result
}

func (d *domainDiff) getChanged() []Subdomain {
	result := make([]Subdomain, 0)

	for _, found := range d.found {

		existing, ok := d.existingMap[found.Name]

		if ok {

			if existing.Cname != found.CName {

				changed := Subdomain{
					Domain: found.Domain,
					Name:   found.Name,
					Cname:  found.CName,
					IsNew:  true,
				}

				result = append(result, changed)
			}
		}
	}

	return result
}

func (d *domainDiff) getDeleted() []Subdomain {
	result := make([]Subdomain, 0)

	for _, existing := range d.existing {

		_, ok := d.foundMap[existing.Name]

		if !ok {
			result = append(result, existing)
		}
	}

	return result
}

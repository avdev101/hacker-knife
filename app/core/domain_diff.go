package core

type domainDiff struct {
	existing []Subdomain
	found    []SubdomainFindItem
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

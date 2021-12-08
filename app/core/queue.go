package core

type ParseSubDomainTask struct {
	DomainName string
}

type ParseIpTask struct {
	DomainName string
}

type ParsePortTask struct {
	IP string
}

type FindServiceTask struct {
	IP string
}

type FindPathTask struct {
	DomainName string
	PortNumber int
	IsHttps    bool
}

type DomainQueue interface {
	PutSubdomainEnumerate(d Domain) error
	TakeSubdomainEnumerate() (Domain, error)
}

package core

type FidnSubDomainTask struct {
	DomainName string
}

type GetIpTask struct {
	DomainName string
}

type FindPortTask struct {
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

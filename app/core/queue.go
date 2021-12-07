package core

type DomainQueue interface {
	PutSubdomainEnumerate(d Domain) error
	TakeSubdomainEnumerate() (Domain, error)
}

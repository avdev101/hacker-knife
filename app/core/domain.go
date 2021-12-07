package core

type Domain struct {
	Name string
}

type Subdomain struct {
	Domain string
	Name   string
	Cname  string
	is_new bool
}

type DomainRepo interface {
	GetList() ([]Domain, error)
	Create(d Domain) error
}

type SubdomainRepo interface {
	GetList(domainName string) ([]Subdomain, error)
	Delete(names []string) error
	CreateBatch(domains []Subdomain) error
	UpdateBatch(domains []Subdomain) error
}

type SubdomainFinder interface {
	Enumerate(domain string) ([]string, error)
}

// DomainCreate Service
type DomainCreateService struct {
	domainRepo DomainRepo
}

func (s *DomainCreateService) Create(d Domain) error {
	if err := s.domainRepo.Create(d); err != nil {
		return err
	}

	return nil
}

// DomainEnumerateService
type DomainEnumerateService struct {
	subdomainRepo SubdomainRepo
	finder        SubdomainFinder
}

func (s *DomainEnumerateService) Enumerate(domain Domain) error {
	// enumerate

	// get existing

	// get deleted
	// remove deleted

	// get changed
	// remove changed

	// create new
	// create changed

	// putIpEnumerate
	// putPortEnumerate

	return nil
}

package core

type Domain struct {
	Name string
}

type Subdomain struct {
	ParentName string
	Name       string
	Cname      string
	IsNew      bool
}

type DomainRepo interface {
	GetList() ([]Domain, error)
	Create(d Domain) error
}

type SubdomainRepo interface {
	Get(name string) (Subdomain, error)
	GetList(domainName string) ([]Subdomain, error)
	DeleteBatch(domains []Subdomain) error
	CreateBatch(domains []Subdomain) error
	UpdateBatch(domains []Subdomain) error
}

// DomainCreate Service
type DomainService struct {
	domainRepo DomainRepo
	taskQeue   TaskQueue
}

func (s *DomainService) Create(d Domain) error {
	if err := s.domainRepo.Create(d); err != nil {
		return err
	}

	task := CollectSubDomainCommand{DomainName: d.Name}
	s.taskQeue.CollectSubdomain(task)

	return nil
}

func (s *DomainService) GetList() ([]Domain, error) {
	domains, err := s.domainRepo.GetList()

	if err != nil {
		return nil, err
	}

	return domains, nil
}

package core

type Domain struct {
	Name string
}

type Subdomain struct {
	Domain string
	Name   string
	Cname  string
	IsNew  bool
}

type DomainRepo interface {
	GetList() ([]Domain, error)
	Create(d Domain) error
}

type SubdomainRepo interface {
	GetList(domainName string) ([]Subdomain, error)
	DeleteBatch(domains []Subdomain) error
	CreateBatch(domains []Subdomain) error
	UpdateBatch(domains []Subdomain) error
}

type SubdomainFindItem struct {
	Name  string
	CName string
}

type SubdomainFinder interface {
	Enumerate(domain string) ([]SubdomainFindItem, error)
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

	task := FidnSubDomainTask{DomainName: d.Name}
	s.taskQeue.FindSubdomain(task)

	return nil
}

func (s *DomainService) GetList() ([]Domain, error) {
	domains, err := s.domainRepo.GetList()

	if err != nil {
		return nil, err
	}

	return domains, nil
}

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

// DomainEnumerateService
type DomainEnumerateService struct {
	subdomainRepo SubdomainRepo
	finder        SubdomainFinder
	taskQeue      TaskQueue
}

func (s *DomainEnumerateService) Enumerate(domainName string, stopPropagate bool) error {

	existing, err := s.subdomainRepo.GetList(domainName)
	if err != nil {
		return err
	}

	domains, err := s.finder.Enumerate(domainName)

	if err != nil {
		return err
	}

	found := make([]SubdomainFindItem, 0)

	for _, d := range domains {

		found = append(found, d)

		if !stopPropagate {
			portTask := FindPortTask{DomainName: d.Name}
			s.taskQeue.FindPort(portTask)

			ipTask := GetIpTask{DomainName: d.Name}
			s.taskQeue.GetIp(ipTask)
		}

	}

	diff := domainDiff{existing, found}
	s.subdomainRepo.CreateBatch(diff.getNew())
	s.subdomainRepo.UpdateBatch(diff.getChanged())
	s.subdomainRepo.DeleteBatch(diff.getDeleted())

	return nil
}

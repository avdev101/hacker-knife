package core

type SubdomainFindItem struct {
	Domain string
	Name   string
	CName  string
}

type SubdomainFinder interface {
	Enumerate(domain string) ([]SubdomainFindItem, error)
}

type DomainEnumerateService struct {
	SubdomainRepo SubdomainRepo
	Finder        SubdomainFinder
	TaskQeue      TaskQueue
}

func (s *DomainEnumerateService) Enumerate(domainName string, stopPropagate bool) error {

	existing, err := s.SubdomainRepo.GetList(domainName)
	if err != nil {
		return err
	}

	domains, err := s.Finder.Enumerate(domainName)

	if err != nil {
		return err
	}

	found := make([]SubdomainFindItem, 0)

	for _, d := range domains {

		found = append(found, d)

		if !stopPropagate {
			portTask := FindPortTask{DomainName: d.Name}
			s.TaskQeue.FindPort(portTask)

			ipTask := GetIpTask{DomainName: d.Name}
			s.TaskQeue.GetIp(ipTask)
		}

	}

	diff := newDomainDiff(existing, found)
	s.SubdomainRepo.CreateBatch(diff.getNew())
	s.SubdomainRepo.UpdateBatch(diff.getChanged())
	s.SubdomainRepo.DeleteBatch(diff.getDeleted())

	return nil
}

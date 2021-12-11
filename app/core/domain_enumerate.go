package core

type SubdomainFinder interface {
	Enumerate(domain string) ([]SubdomainFindItem, error)
}

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

package core

type SubdomainCollectItem struct {
	Domain string
	Name   string
	CName  string
}

type SubdomainCollector interface {
	Collect(domain string) ([]SubdomainCollectItem, error)
}

type DomainCollectService struct {
	SubdomainRepo SubdomainRepo
	Finder        SubdomainCollector
	TaskQeue      TaskQueue
}

func (s *DomainCollectService) Collect(domainName string, stopPropagate bool) error {

	existing, err := s.SubdomainRepo.GetListByParent(domainName)
	if err != nil {
		return err
	}

	domains, err := s.Finder.Collect(domainName)

	if err != nil {
		return err
	}

	found := make([]SubdomainCollectItem, 0)

	for _, d := range domains {

		found = append(found, d)

		if !stopPropagate {
			portTask := CollectPortCommand{DomainName: d.Name}
			s.TaskQeue.CollectPort(portTask)

			ipTask := CollectIPCommand{DomainName: d.Name}
			s.TaskQeue.CollectIP(ipTask)
		}

	}

	diff := newDomainDiff(existing, found)
	s.SubdomainRepo.CreateBatch(diff.getNew())
	s.SubdomainRepo.UpdateBatch(diff.getChanged())
	s.SubdomainRepo.DeleteBatch(diff.getDeleted())

	return nil
}

package core

type IPCollectService struct {
	IPRepo        IPRepo
	DomainRepo    DomainRepo
	SubdomainRepo SubdomainRepo
	TaskQueue     TaskQueue
	IPCollector   IPCollector
}

func (s *IPCollectService) Collect(domain string, stopPropagate bool) error {

	d, err := s.SubdomainRepo.Get(domain)
	if err != nil {
		return err
	}

	existing, err := s.IPRepo.GetBySubdomain(domain)
	if err != nil {
		return err
	}

	collected, err := s.IPCollector.Collect(domain)
	if err != nil {
		return err
	}

	err = s.saveDiff(d, existing, collected)
	if err != nil {
		return err
	}

	if !stopPropagate {
		s.propagateTask(collected)
	}

	return nil
}

func (s *IPCollectService) createNewIPList(d Subdomain, collectedIPs []IPCollectItem) []IP {
	result := make([]IP, len(collectedIPs))

	for i, found := range collectedIPs {
		ip := IP{
			ParentDomain: d.ParentName,
			Domain:       d.Name,
			Addr:         found.Addr,
		}
		result[i] = ip
	}

	return result
}

func (s *IPCollectService) saveDiff(d Subdomain, existing []IP, collected []IPCollectItem) error {

	diff := IPDiff{existing, collected}

	newIPList := s.createNewIPList(d, diff.getNew())
	s.IPRepo.CreateBatch(newIPList)
	s.IPRepo.DeleteBatch(diff.getDeleted())

	return nil
}

func (s *IPCollectService) propagateTask(collected []IPCollectItem) error {

	for _, ip := range collected {
		t := CollectServiceCommand{IP: ip.Addr}
		s.TaskQueue.CollectService(t)

	}

	return nil
}

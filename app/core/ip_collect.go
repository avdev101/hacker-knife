package core

type IPCollectService struct {
	IPRepo      IPRepo
	DomainRepo  DomainRepo
	TaskQeue    TaskQueue
	IPCollector IPCollector
}

func (s *IPCollectService) Collect(domain string, stopPropagate bool) error {

	existing, err := s.IPRepo.GetBySubdomain(domain)
	if err != nil {
		return err
	}

	collected, err := s.IPCollector.Collect(domain)
	if err != nil {
		return err
	}

	err = s.saveDiff(existing, collected)
	if err != nil {
		return err
	}

	if !stopPropagate {
		s.propagateTask(collected)
	}

	return nil
}

func (s *IPCollectService) saveDiff(existing []IP, collected []IPCollectItem) error {

	diff := IPDiff{existing, collected}

	s.IPRepo.CreateBatch(diff.getNew())
	s.IPRepo.DeleteBatch(diff.getDeleted())

	return nil
}

func (s *IPCollectService) propagateTask(collected []IPCollectItem) error {

	for _, ip := range collected {
		t := CollectServiceCommand{IP: ip.Addr}
		s.TaskQeue.CollectService(t)

	}

	return nil
}

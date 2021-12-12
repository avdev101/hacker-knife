package core

type IPCollectService struct {
	IPRepo      IPRepo
	DomainRepo  DomainRepo
	TaskQeue    TaskQueue
	IPCollector IPCollector
}

func (s *IPCollectService) Collect(domain string) error {
	return nil
}

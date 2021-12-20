package core

type Service struct {
	Domain      string
	IPAddr      string
	PortNumber  int
	ServiceType string
	IsNew       bool
}

type ServiceRepo interface {
	GetListByDomain(name string) ([]Service, error)
	GetListByIP(addr string) ([]Service, error)
	CreateBatch([]Service) error
	DeleteBatch([]Service) error
}

type ServiceCollectItem struct {
	IPAddr      string
	ServiceType string
	PortNumber  int
}

type ServiceCollector interface {
	Collect(addr string) ([]ServiceCollectItem, error)
}

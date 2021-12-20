package core

type Service struct {
	Domain      string
	IPAddr      string
	PortNumber  int
	ServiceType string
	IsNew       bool
}

type ServiceRepo interface {
	GetListByIP(addr string) ([]Service, error)
	CreateBatch([]Service) error
	DeleteBatch([]Service) error
}

type ServiceCollectItem struct {
	IPAddr      string
	ServiceType string
	PortNumber  string
}

type ServiceCollector interface {
	Collect(addr string) ([]ServiceCollectItem, error)
}

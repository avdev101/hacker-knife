package core

type IP struct {
	ParentDomain string
	Domain       string
	Addr         string
	IsNew        bool
}

type IPRepo interface {
	GetByDomain(name string) ([]IP, error)
	GetBySubdomain(name string) ([]IP, error)
	CreateBatch(ips []IP) error
	DeleteBatch(ips []IP) error
}

type IPCollectItem struct {
	Domain string
	Addr   string
}

type IPCollector interface {
	Collect(domain string) ([]IPCollectItem, error)
}

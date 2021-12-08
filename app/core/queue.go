package core

type FidnSubDomainTask struct {
	DomainName string
}

type GetIpTask struct {
	DomainName string
}

type FindPortTask struct {
	IP string
}

type FindServiceTask struct {
	IP string
}

type FindPathTask struct {
	DomainName string
	PortNumber int
	IsHttps    bool
}

type MakeShotTask struct {
	DomainName string
	PortNumber int
	Path       int
	IsHttps    bool
}

type TaskQueue interface {
	FindSubdomain(t FidnSubDomainTask) error
	GetIp(t GetIpTask) error
	FindPort(t FindPortTask) error
	FindService(t FindServiceTask) error
	FindPath(t FindPathTask) error
	MakeShot(t MakeShotTask) error
}

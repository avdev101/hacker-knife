package core

type TaskMeta struct {
	Propagate bool
}

type FidnSubDomainTask struct {
	DomainName string
	TaskMeta
}

type GetIpTask struct {
	DomainName string
	TaskMeta
}

type FindPortTask struct {
	IP string
	TaskMeta
}

type FindServiceTask struct {
	IP string
	TaskMeta
}

type FindPathTask struct {
	DomainName string
	PortNumber int
	IsHttps    bool
	TaskMeta
}

type MakeShotTask struct {
	DomainName string
	PortNumber int
	Path       int
	IsHttps    bool
	TaskMeta
}

type TaskQueue interface {
	FindSubdomain(t FidnSubDomainTask) error
	GetIp(t GetIpTask) error
	FindPort(t FindPortTask) error
	FindService(t FindServiceTask) error
	FindPath(t FindPathTask) error
	MakeShot(t MakeShotTask) error
}

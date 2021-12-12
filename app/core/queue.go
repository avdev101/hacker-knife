package core

type TaskMeta struct {
	Propagate bool
}

type CollectSubDomainCommand struct {
	DomainName string
	TaskMeta
}

type CollectSubdomainTask struct {
	Ack  func()
	Nack func()
	Data CollectSubDomainCommand
}

type CollectIPCommand struct {
	DomainName string
	TaskMeta
}

type CollectIPTask struct {
	Ack  func()
	Nack func()
	Data CollectIPCommand
}

type CollectPortCommand struct {
	DomainName string
	TaskMeta
}

type CollectServiceCommand struct {
	IP string
	TaskMeta
}

type CollectPathCommand struct {
	DomainName string
	PortNumber int
	IsHttps    bool
	TaskMeta
}

type CollectShotCommand struct {
	DomainName string
	PortNumber int
	Path       int
	IsHttps    bool
	TaskMeta
}

type TaskQueue interface {
	CollectSubdomain(t CollectSubDomainCommand) error
	TakeCollectSubdomain() (CollectSubdomainTask, error)

	CollectIP(t CollectIPCommand) error
	TakeCollectIP() (CollectIPTask, error)

	CollectPort(t CollectPortCommand) error
	CollectService(t CollectServiceCommand) error
	CollectPath(t CollectPathCommand) error
	CollectShot(t CollectShotCommand) error
}

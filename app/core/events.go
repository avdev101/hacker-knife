package core

type DomainEvent struct {
	Name string
}

type SubdomainEvent struct {
	Name string
}

type IPEvent struct {
	IP string
}

type PortEvent struct {
	DomainName string
	PortNumber int
	IsHttps    bool
}

type PathEvent struct {
	DomainName string
	Path       string
	PortNumber int
	IsHttps    bool
}

type ServiceEvent struct {
	IP         string
	PortNumber int
	Type       string
}

type EventEmitter interface {
	DomainAdd(e DomainEvent) error
	SubdomainAdd(e SubdomainEvent) error
	IpAdd(e IPEvent) error
	PortAdd(e PortEvent) error
	PathAdd(e PathEvent) error
}

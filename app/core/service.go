package core

type Service struct {
	ParentDomain string
	Domain       string
	IPAddr       string
	PortNumber   int
	ServiceType  string
	IsNew        bool
}

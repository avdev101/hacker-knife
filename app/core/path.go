package core

type Path struct {
	ParentDomain string
	Subdomain    string
	PortNumber   int
	IsHTTPS      bool
	Path         string
	IsNew        bool
}

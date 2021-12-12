package core

type Path struct {
	ParentDomain string
	Domain       string
	PortNumber   int
	IsHTTPS      bool
	Path         string
	IsNew        bool
}

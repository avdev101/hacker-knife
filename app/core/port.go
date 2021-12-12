package core

type Port struct {
	ParentDomain string
	Domain       string
	Number       int
	IsHTTPS      bool
	IsNew        bool
}

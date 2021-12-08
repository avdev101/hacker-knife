package core

type Port struct {
	ParentDomain string
	Subdomain    string
	Number       int
	IsHttps      bool
	IsNew        bool
}

package core

type Service struct {
	ParentDomain string
	Subdomain    string
	IP           string
	ServiceType  string
	IsNew        bool
}

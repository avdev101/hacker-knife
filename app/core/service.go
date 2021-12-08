package core

type Service struct {
	ParentDomain string
	Subdomain    string
	IP           string
	Port         int
	ServiceType  string
	IsNew        bool
}

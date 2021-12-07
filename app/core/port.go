package core

type Port struct {
	ParentDomain string
	Subdomain    string
	Number       int
	Is_https     bool
	Is_new       bool
}

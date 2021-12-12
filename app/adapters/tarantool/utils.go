package tarantool

import "github.com/eremeevdev/hacker-knife/core"

func tuplesToSubdomains(tuples [][]interface{}) []core.Subdomain {
	result := make([]core.Subdomain, 0)

	for _, tuple := range tuples {
		item := core.Subdomain{
			Domain: tuple[0].(string),
			Name:   tuple[1].(string),
			Cname:  tuple[2].(string),
			IsNew:  tuple[3].(bool),
		}
		result = append(result, item)
	}

	return result
}

func subdomainToTuples(domains []core.Subdomain) []interface{} {
	result := make([]interface{}, len(domains))

	for i, d := range domains {
		item := []interface{}{
			d.Domain,
			d.Name,
			d.Cname,
			d.IsNew,
		}
		result[i] = item
	}

	return result
}

func tuplesToDomains(tuples [][]interface{}) []core.Domain {
	result := make([]core.Domain, len(tuples))

	for i, tuple := range tuples {
		item := core.Domain{Name: tuple[0].(string)}
		result[i] = item
	}

	return result
}

func domainToTuple(domain core.Domain) []interface{} {
	return []interface{}{domain.Name}
}

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

package core

type IPDiff struct {
	existing []IP
	found    []IPCollectItem
}

func (d *IPDiff) getExistingMap() map[string]IP {
	result := make(map[string]IP)

	for _, ip := range d.existing {
		result[ip.Addr] = ip
	}

	return result
}

func (d *IPDiff) makeNewIP(foundIP IPCollectItem) IP {
	return IP{
		Domain: foundIP.Domain,
		Addr:   foundIP.Addr,
		IsNew:  true,
	}
}

func (d *IPDiff) getNew() []IP {
	existingMap := d.getExistingMap()

	result := make([]IP, 0)

	for _, foundIP := range d.found {
		_, ok := existingMap[foundIP.Addr]

		if !ok {
			result = append(result, d.makeNewIP(foundIP))
		}
	}

	return result
}

func (d IPDiff) getCollectedMap() map[string]IPCollectItem {
	result := make(map[string]IPCollectItem)

	for _, ip := range d.found {
		result[ip.Addr] = ip
	}

	return result
}

func (d *IPDiff) getDeleted() []IP {
	result := make([]IP, 0)

	collectedMap := d.getCollectedMap()

	for _, ip := range d.existing {
		_, ok := collectedMap[ip.Addr]

		if !ok {
			result = append(result, ip)
		}
	}

	return result
}

package core

type IPDiff struct {
	existing []IP
	found    []IPCollectItem
}

func (d *IPDiff) getNew() []IP {
	return nil
}

func (d *IPDiff) getDeleted() []IP {
	return nil
}

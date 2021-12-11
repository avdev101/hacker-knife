package tarantool

import "github.com/eremeevdev/hacker-knife/core"

type Queue struct {
}

func (q *Queue) FindSubdomain(t core.FidnSubDomainTask) error {
	return nil
}

func (q *Queue) GetIp(t core.GetIpTask) error {
	return nil
}

func (q *Queue) FindPort(t core.FindPortTask) error {
	return nil
}

func (q *Queue) FindService(t core.FindServiceTask) error {
	return nil
}

func (q *Queue) FindPath(t core.FindPathTask) error {
	return nil
}

func (q *Queue) MakeShot(t core.MakeShotTask) error {
	return nil
}

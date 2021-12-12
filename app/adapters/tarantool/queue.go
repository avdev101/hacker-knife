package tarantool

import (
	"encoding/json"
	"fmt"

	"github.com/eremeevdev/hacker-knife/core"
	"github.com/tarantool/go-tarantool"
	"github.com/tarantool/go-tarantool/queue"
)

type Queue struct {
	conn *tarantool.Connection
}

func NewQueue(host string, user string, pass string) (Queue, error) {
	conn, err := tarantool.Connect(host, tarantool.Opts{
		User: user,
		Pass: pass,
	})

	if err != nil {
		return Queue{}, fmt.Errorf("can't create connection: %v", err)
	}

	return Queue{conn}, nil
}

func (q *Queue) FindSubdomain(t core.FidnSubDomainTask) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	que := queue.New(q.conn, "parse_subdomain")
	_, err = que.Put(data)
	return err
}

func (q *Queue) TakeSubdomain() (core.FindSubdomainConsume, error) {
	que := queue.New(q.conn, "parse_subdomain")

	var res core.FindSubdomainConsume

	task, err := que.Take()
	if err != nil {
		return res, err
	}

	res.Ack = func() {
		task.Ack()
	}

	res.Nack = func() {
		task.Release()
	}

	bData := []byte(task.Data().(string))

	err = json.Unmarshal(bData, &res.Data)
	if err != nil {
		return res, err
	}

	return res, nil
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

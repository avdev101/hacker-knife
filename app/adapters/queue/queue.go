package queue

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

func (q *Queue) CollectSubdomain(t core.CollectSubDomainCommand) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	que := queue.New(q.conn, "parse_subdomain")
	_, err = que.Put(data)
	return err
}

func (q *Queue) TakeCollectSubdomain() (core.CollectSubdomainTask, error) {
	que := queue.New(q.conn, "parse_subdomain")

	var res core.CollectSubdomainTask

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

func (q *Queue) CollectIP(t core.CollectIPCommand) error {
	return nil
}

func (q *Queue) CollectPort(t core.CollectPortCommand) error {
	return nil
}

func (q *Queue) CollectService(t core.CollectServiceCommand) error {
	return nil
}

func (q *Queue) CollectPath(t core.CollectPathCommand) error {
	return nil
}

func (q *Queue) CollectShot(t core.CollectShotCommand) error {
	return nil
}

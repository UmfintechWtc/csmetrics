package client

import (
	"errors"
	"fmt"
	"sync"

	"github.com/abrander/go-supervisord"
)

var (
	processes  = &Supervisor{}
	SpOnce     sync.Once
	err        error
	supervisor *supervisord.Client
)

type Supervisor struct {
	AllProcessInfo []string
}

func NewSupervisor(ip string, port int, version, state string) (*Supervisor, error) {
	SpOnce.Do(func() {
		supervisor, err = supervisord.NewClient(
			fmt.Sprintf("http://%s:%d/RPC2", ip, port),
		)
	})
	if err != nil {
		return nil, err
	}

	v, err := supervisor.GetAPIVersion()
	if err != nil {
		return nil, err
	}

	s, err := supervisor.GetState()
	if err != nil {
		return nil, err
	}

	if v != version && string(s.Name) != state {
		return nil, errors.New(
			fmt.Sprintf(
				"invalid version [%s] or state []%s", v, string(s.Name),
			),
		)
	}

	allProcess, err := supervisor.GetAllProcessInfo()
	if err != nil {
		return nil, err
	}

	for _, process := range allProcess {
		processes.AllProcessInfo = append(processes.AllProcessInfo, process.Name)
	}
	return processes, nil
}

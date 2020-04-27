package integration

import (
	"github.com/docker/docker/client"
)

type container struct {
	Image string
}

type suite struct {
	cli        *client.Client
	containers []*container
}

func (s *suite) addContainer(c *container) {
	s.containers = append(s.containers, c)
}

func (s *suite) SetUp() error {
	return nil
}

func (s *suite) Run() error {
	return nil
}

func (s *suite) Down() {
}

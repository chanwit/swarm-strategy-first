package main

import (
	"github.com/docker/swarm/scheduler/strategy/plugin"
	"github.com/samalba/dockerclient"
)

type FirstOnlyStrategy struct{}

func (s *FirstOnlyStrategy) Name() string {
	return "first"
}

func (s *FirstOnlyStrategy) Initialize() error {
	return nil
}

func (s *FirstOnlyStrategy) PlaceContainer(config *dockerclient.ContainerConfig, nodes []*plugin.Node) (*plugin.Node, error) {
	return nodes[0], nil
}

func main() {
	plugin.Run(new(FirstOnlyStrategy))
}

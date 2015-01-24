package main

import (
	"github.com/chanwit/tentacle"
	"github.com/samalba/dockerclient"
)

type FirstOnlyStrategy struct{}

func (s *FirstOnlyStrategy) Initialize() error {
	return nil
}

func (s *FirstOnlyStrategy) PlaceContainer(config *dockerclient.ContainerConfig, nodes []*tentacle.Node) (*tentacle.Node, error) {
	return nodes[0], nil
}

func main() {
	tentacle.Run(&FirstOnlyStrategy{})
}

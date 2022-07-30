package balancer

import (
	"github.com/balchua/consistent-demo/pkg/config"
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

type clusterMember string

func (m clusterMember) String() string {
	return string(m)
}

type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

type ConsistentHashLoadBalancer struct {
	c *consistent.Consistent
}

func NewBalancer(clusterConfig config.ClusterConfig) *ConsistentHashLoadBalancer {
	// Create a new consistent instance
	cfg := consistent.Config{
		PartitionCount:    30,
		ReplicationFactor: 20,
		Load:              1.25,
		Hasher:            hasher{},
	}
	b := &ConsistentHashLoadBalancer{
		c: consistent.New(nil, cfg),
	}

	for _, node := range clusterConfig.Infra.Nodes {
		b.c.Add(clusterMember(node.Name))
	}

	return b
}

func (b *ConsistentHashLoadBalancer) Pick(name string) consistent.Member {
	return b.c.LocateKey([]byte(name))
}

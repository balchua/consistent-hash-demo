package balancer

import (
	"github.com/balchua/consistent-demo/pkg/config"
	"github.com/balchua/consistent-demo/pkg/logging"
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	cmap "github.com/orcaman/concurrent-map/v2"
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
	c           *consistent.Consistent
	nodeHistory cmap.ConcurrentMap[consistent.Member]
}

func NewBalancer(clusterConfig config.ClusterConfig) *ConsistentHashLoadBalancer {
	// Create a new consistent instance
	cfg := consistent.Config{
		PartitionCount:    50,
		ReplicationFactor: 20,
		Load:              1.25,
		Hasher:            hasher{},
	}
	b := &ConsistentHashLoadBalancer{
		c:           consistent.New(nil, cfg),
		nodeHistory: cmap.New[consistent.Member](),
	}

	for _, node := range clusterConfig.Infra.Nodes {
		b.c.Add(clusterMember(node.Name))
	}

	go b.checkHealth(clusterConfig)

	return b
}

func (b *ConsistentHashLoadBalancer) Pick(key string) (previousMember consistent.Member, currentMember consistent.Member) {
	currentNode := b.c.LocateKey([]byte(key))
	previous, _ := b.nodeHistory.Get(key)
	if previous == nil {
		previous = clusterMember("")
	}
	if previous != currentNode {
		b.nodeHistory.Set(key, currentNode)
		logging.Infof("Previous Member: %s , Current Member: %s", previous, currentNode)
	}

	return previous, currentNode
}

func (b *ConsistentHashLoadBalancer) AddNode(nodeName string) {
	b.c.Add(clusterMember(nodeName))
}

func (b *ConsistentHashLoadBalancer) RemoveNode(nodeName string) {
	b.c.Remove(nodeName)
}

func (b *ConsistentHashLoadBalancer) List() map[string]int {
	owners := make(map[string]int)
	for partID := 0; partID < 50; partID++ {
		owner := b.c.GetPartitionOwner(partID)
		owners[owner.String()]++
	}

	return owners
}

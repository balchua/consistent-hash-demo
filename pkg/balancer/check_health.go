package balancer

import (
	"net"
	"time"

	"github.com/balchua/consistent-demo/pkg/config"
	"github.com/balchua/consistent-demo/pkg/logging"
)

func (b *ConsistentHashLoadBalancer) checkHealth(clusterConfig config.ClusterConfig) {
	for {
		for _, member := range clusterConfig.Infra.Nodes {
			timeout := time.Duration(1 * time.Second)
			_, err := net.DialTimeout("tcp", member.Name, timeout)
			if err != nil {
				logging.Infof("member %s is not reachable, removing from member list", member.Name)
				b.c.Remove(member.Name)
			} else {
				b.c.Add(clusterMember(member.Name))
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}

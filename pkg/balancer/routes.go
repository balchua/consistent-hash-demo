package balancer

import (
	"fmt"

	"github.com/balchua/consistent-demo/pkg/logging"
	"github.com/gofiber/fiber/v2"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type lbServer struct {
	chlb        *ConsistentHashLoadBalancer
	nodeHistory cmap.ConcurrentMap[string]
}

func NewServer(chlb *ConsistentHashLoadBalancer) *lbServer {

	return &lbServer{
		chlb:        chlb,
		nodeHistory: cmap.New[string](),
	}
}

func (lb *lbServer) Calculate(c *fiber.Ctx) error {
	key := c.Params("key")
	node := lb.chlb.Pick(key)
	history, _ := lb.nodeHistory.Get(key)
	proxyRequest := proxyRequest{
		previousNode: node.String(),
		url:          fmt.Sprintf("http://%s/calculate", node),
	}
	if history != node.String() {
		lb.nodeHistory.Set(key, node.String())
		logging.Infof("Node: %s is added to history", node.String())
		proxyRequest.previousNode = history
	}
	proxy(proxyRequest)
	return c.SendString(fmt.Sprintf("node: %s ", node))
}

func (lb *lbServer) RemoveNode(c *fiber.Ctx) error {
	node := c.FormValue("node")
	lb.chlb.c.Remove(node)
	return c.SendString(fmt.Sprintf("node: %s is removed", node))
}

func (lb *lbServer) AddNode(c *fiber.Ctx) error {
	node := c.FormValue("node")
	lb.chlb.c.Add(clusterMember(node))
	return c.SendString(fmt.Sprintf("node: %s is added", node))
}

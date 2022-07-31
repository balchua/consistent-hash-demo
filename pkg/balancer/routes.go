package balancer

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type clusterNode struct {
	PreviousNode string `json:"previousNode"`
	CurrentNode  string `json:"currentNode"`
}

type lbServer struct {
	chlb *ConsistentHashLoadBalancer
}

func NewServer(chlb *ConsistentHashLoadBalancer) *lbServer {

	return &lbServer{
		chlb: chlb,
	}
}

func (lb *lbServer) Pick(c *fiber.Ctx) error {
	key := c.Params("key")
	previousNode, currentNode := lb.chlb.Pick(key)

	cn := clusterNode{
		PreviousNode: previousNode.String(),
		CurrentNode:  currentNode.String(),
	}

	responseInJson, err := json.Marshal(cn)
	if err != nil {
		return err
	}
	return c.SendString(string(responseInJson))
}

func (lb *lbServer) RemoveNode(c *fiber.Ctx) error {
	node := c.FormValue("node")
	lb.chlb.RemoveNode(node)
	return c.SendString(fmt.Sprintf("node: %s is removed", node))
}

func (lb *lbServer) AddNode(c *fiber.Ctx) error {
	node := c.FormValue("node")
	lb.chlb.AddNode(node)
	return c.SendString(fmt.Sprintf("node: %s is added", node))
}

func (lb *lbServer) List(c *fiber.Ctx) error {

	owners := lb.chlb.List()
	jsonStr, err := json.Marshal(owners)
	if err != nil {
		return err
	}
	return c.SendString(string(jsonStr))
}

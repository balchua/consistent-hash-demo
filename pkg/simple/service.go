package simple

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (calc *service) Do(c *fiber.Ctx) error {
	// simulate insert to db
	time.Sleep(time.Duration(20 * time.Millisecond))
	node := c.Get("previousNode")
	currentHost := c.Context().URI().Host()
	if node != string(currentHost) {
		// simulate retrieve data from db
		time.Sleep(time.Duration(100 * time.Millisecond))
		return c.SendString(fmt.Sprintf("node %s is not the same as current host %s, clearing cache", node, currentHost))
	}
	// simulate calculate data
	time.Sleep(time.Duration(20 * time.Millisecond))
	return c.SendString("same same")
}

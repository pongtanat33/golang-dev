package servers

import (
	"fmt"

	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/api/v1")

	s.App.Use(func(c *fiber.Ctx) error {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "error, end point not found", nil)
	})
	return nil
}

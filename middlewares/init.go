package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	db *sqlx.DB
)

func NewMiddlewares(app fiber.Router, d *sqlx.DB) {
	db = d
	app.Use(
		realIP,
		recover.New(),
		helmet.New(),
		cors.New(),
		requestid.New(),
		compress.New(),
	)

	limiterConfig := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        100,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return fiber.ErrTooManyRequests
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Locals(constants.RealIPKey).(string)
		},
	}

	app.Use(limiter.New(limiterConfig))

	if envValue := viper.GetString("ENV.LOCAL"); envValue == "1" || envValue == "2" {
		app.Use(NewLoggerDevelopment)
	} else {
		app.Use(NewLoggerProduction)
	}
}

func realIP(c *fiber.Ctx) error {
	var realIP string

	if raw := c.Get(constants.HeaderXOriginalForwardedFor); raw != "" {
		ips := strings.Split(strings.ReplaceAll(raw, " ", ""), ",")
		if len(ips) > 0 {
			realIP = ips[0]
		}
	}

	if realIP == "" {
		realIP = c.IP()
	}

	c.Locals(constants.RealIPKey, realIP)
	return c.Next()
}

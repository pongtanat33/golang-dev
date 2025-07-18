package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/bytebufferpool"
)

type Logger struct {
	Timestamp string `json:"timestamp"`
	ReqeustID string `json:"request_id"`
	Status    int    `json:"status"`
	Latency   string `json:"latency"`
	IP        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
}

var (
	timeFormat   = "2006-01-02T15:04:05.000Z0700"
	timeInterval = 500 * time.Millisecond
)

func NewLoggerDevelopment(c *fiber.Ctx) error {
	var timestamp atomic.Value
	timestamp.Store(time.Now().UTC().In(time.Local).Format(timeFormat))
	go func() {
		for {
			time.Sleep(timeInterval)
			timestamp.Store(time.Now().UTC().In(time.Local).Format(timeFormat))
		}
	}()

	var start, stop time.Time
	start = time.Now().UTC()
	chainErr := c.Next()
	if chainErr != nil {
		return chainErr
	}
	stop = time.Now().UTC()

	buf := bytebufferpool.Get()
	_, _ = buf.WriteString(fmt.Sprintf("%s | %s | %3d | %7v | %15s | %-7s | %s\n",
		timestamp.Load().(string),
		c.Locals("requestid"),
		c.Response().StatusCode(),
		stop.Sub(start).Round(time.Millisecond),
		c.Locals(constants.RealIPKey),
		c.Method(),
		c.Path(),
	))

	_, _ = os.Stdout.Write(buf.Bytes())
	bytebufferpool.Put(buf)

	return nil
}

func NewLoggerProduction(c *fiber.Ctx) error {
	var timestamp atomic.Value
	timestamp.Store(time.Now().UTC().In(time.Local).Format(timeFormat))
	go func() {
		for {
			time.Sleep(timeInterval)
			timestamp.Store(time.Now().UTC().In(time.Local).Format(timeFormat))
		}
	}()

	var start, stop time.Time
	start = time.Now().UTC()
	chainErr := c.Next()
	if chainErr != nil {
		return chainErr
	}
	stop = time.Now().UTC()

	buf := bytebufferpool.Get()

	logger := Logger{
		Timestamp: timestamp.Load().(string),
		ReqeustID: c.Locals("requestid").(string),
		Status:    c.Response().StatusCode(),
		Latency:   stop.Sub(start).Round(time.Millisecond).String(),
		IP:        c.Locals(constants.RealIPKey).(string),
		Method:    c.Method(),
		Path:      c.Path(),
	}

	logStr, _ := json.Marshal(logger)

	_, _ = buf.WriteString(string(logStr) + "\n")
	_, _ = os.Stdout.Write(buf.Bytes())
	bytebufferpool.Put(buf)

	return nil
}

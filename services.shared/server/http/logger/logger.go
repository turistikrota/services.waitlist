package logger

import (
	"opensource.turistikrota.com/shared/formats"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func New() fiber.Handler {
	// return fiber logger with logrus
	return logger.New(logger.Config{
		Format:     "${method} ${status} ${path}: ${time} - ${latency} - ${ip}\n",
		TimeFormat: formats.DateYYYYMMDDHHMMSS,
		TimeZone:   "Local",
		Output:     logrus.StandardLogger().Out,
	})
}

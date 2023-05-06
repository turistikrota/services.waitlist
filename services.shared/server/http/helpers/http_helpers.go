package http_helpers

import (
	"time"

	"opensource.turistikrota.com/shared/server/http/parser"
	"opensource.turistikrota.com/shared/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/i18np"
)

type Helpers interface {
	ParseBody(c *fiber.Ctx, d interface{})
	ParseParams(c *fiber.Ctx, d interface{})
	ParseQuery(c *fiber.Ctx, d interface{})
	WrapWithTimeout(fn fiber.Handler) fiber.Handler
}

type helpers struct {
	i18n      i18np.I18n
	validator validator.Validator
}

type Config struct {
	I18n      i18np.I18n
	Validator validator.Validator
}

func New(config Config) Helpers {
	return &helpers{
		i18n:      config.I18n,
		validator: config.Validator,
	}
}

func (h *helpers) ParseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, h.i18n, d)
}

func (h *helpers) ParseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, h.i18n, d)
}

func (h *helpers) ParseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, h.i18n, d)
}

func (h *helpers) WrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.New(fn, 10*time.Second)
}

package http

import (
	"context"
	"time"

	"opensource.turistikrota.com/shared/server/http/parser"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"opensource.turistikrota.com/shared/validator"
	"opensource.turistikrota.com/waitlist/src/app"
	"opensource.turistikrota.com/waitlist/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/i18np"
)

type Server struct {
	app         app.Application
	i18n        i18np.I18n
	validator   validator.Validator
	ctx         context.Context
	httpHeaders config.HttpHeaders
}

type Config struct {
	App         app.Application
	I18n        i18np.I18n
	Validator   validator.Validator
	Context     context.Context
	HttpHeaders config.HttpHeaders
}

func New(config Config) Server {
	return Server{
		app:         config.App,
		i18n:        config.I18n,
		validator:   config.Validator,
		ctx:         config.Context,
		httpHeaders: config.HttpHeaders,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Use(h.cors(), h.rateLimit())
	router.Post("/", h.wrapWithTimeout(h.Join))
	router.Get("/:token", h.wrapWithTimeout(h.Leave))
	return router
}

func (h Server) parseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, h.i18n, d)
}

func (h Server) parseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, h.i18n, d)
}

func (h Server) parseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, h.i18n, d)
}

func (h Server) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}

func (h Server) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.httpHeaders.AllowedOrigins,
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
	})
}

func (h Server) rateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10,
		Expiration: 5 * time.Minute,
	})
}

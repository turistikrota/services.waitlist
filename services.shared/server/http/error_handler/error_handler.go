package error_handler

import (
	i18nHttp "opensource.turistikrota.com/shared/server/http/i18n"
	"opensource.turistikrota.com/shared/server/http/result"

	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
)

type Config struct {
	DfMsgKey string
	I18n     *i18np.I18n
}

func New(cfg Config) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Status).JSON(e)
		}
		if e, ok := err.(*result.DetailResult); ok {
			return c.Status(e.Status).JSON(e.Detail)
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if cfg.DfMsgKey != "" {
			l, a := i18nHttp.GetLanguagesInContext(*cfg.I18n, c)
			return c.Status(code).JSON(result.Error(cfg.I18n.Translate(cfg.DfMsgKey, l, a), code))
		}
		err = c.Status(code).JSON(result.Error(err.Error(), code))
		return err
	}
}

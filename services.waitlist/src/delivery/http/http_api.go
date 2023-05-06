package http

import (
	httpI18n "opensource.turistikrota.com/shared/server/http/i18n"
	"opensource.turistikrota.com/shared/server/http/result"
	"opensource.turistikrota.com/waitlist/src/delivery/http/dto"
	"github.com/gofiber/fiber/v2"
)

func (h Server) Join(ctx *fiber.Ctx) error {
	d := dto.Request.Join()
	h.parseBody(ctx, d)
	lang, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	_, err := h.app.Commands.Join.Handle(ctx.UserContext(), d.ToCommand(lang))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.WaitlistJoined)
}

func (h Server) Leave(ctx *fiber.Ctx) error {
	d := dto.Request.Leave()
	h.parseParams(ctx, d)
	lang, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	_, err := h.app.Commands.Leave.Handle(ctx.UserContext(), d.ToCommand(lang))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.WaitlistLeaved)
}

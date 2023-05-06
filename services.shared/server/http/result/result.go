package result

import (
	httpI18n "opensource.turistikrota.com/shared/server/http/i18n"
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
)

type Result struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

type DetailResult struct {
	Result
	Detail any `json:"detail"`
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DetailResult) Error() string {
	return r.Message
}

func Success(m string, c ...int) *Result {
	code := fiber.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func IfSuccess(err *i18np.Error, ctx *fiber.Ctx, i18n i18np.I18n, msg string) *Result {
	l, a := httpI18n.GetLanguagesInContext(i18n, ctx)
	if err != nil {
		return Error(i18n.TranslateFromError(*err, l, a))
	}
	return Success(i18n.Translate(msg, l, a))
}

func IfSuccessParams(err *i18np.Error, ctx *fiber.Ctx, i18n i18np.I18n, msg string) error {
	l, a := httpI18n.GetLanguagesInContext(i18n, ctx)
	if err != nil {
		return ErrorDetail(i18n.TranslateFromErrorDetail(*err, l, a))
	}
	return Success(i18n.Translate(msg, l, a))
}

func Error(m string, c ...int) *Result {
	code := fiber.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func SuccessDetail(m string, d any, c ...int) *DetailResult {
	code := fiber.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}

func IfSuccessDetail(err *i18np.Error, ctx *fiber.Ctx, i18n i18np.I18n, msg string, mapper func() interface{}) error {
	l, a := httpI18n.GetLanguagesInContext(i18n, ctx)
	if err != nil {
		return Error(i18n.TranslateFromError(*err, l, a))
	}
	return SuccessDetail(i18n.Translate(msg, l, a), mapper())
}

func IfSuccessDetailParams(err *i18np.Error, ctx *fiber.Ctx, i18n i18np.I18n, msg string, mapper func() interface{}) error {
	l, a := httpI18n.GetLanguagesInContext(i18n, ctx)
	if err != nil {
		return ErrorDetail(i18n.TranslateFromErrorDetail(*err, l, a))
	}
	return SuccessDetail(i18n.Translate(msg, l, a), mapper())
}

func ErrorDetail(m string, d any, c ...int) *DetailResult {
	code := fiber.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}

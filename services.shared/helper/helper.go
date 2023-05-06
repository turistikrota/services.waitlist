package helper

import (
	"opensource.turistikrota.com/shared/helper/languages"
	notify_helper "opensource.turistikrota.com/shared/helper/notify"
	"opensource.turistikrota.com/shared/helper/parsers"
)

var (
	Notify   notify_helper.Helper = notify_helper.New()
	Language                      = languages.New()
	Parsers  parsers.Parser       = parsers.New()
)

func Fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnIfErrNil[Type any](err error, val Type) (Type, error) {
	if err != nil {
		return val, err
	}
	return val, nil
}

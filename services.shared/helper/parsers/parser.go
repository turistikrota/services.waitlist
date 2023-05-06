package parsers

import (
	"strconv"
	"strings"
)

type Parser interface {
	ParseUsernameAndCode(str string) (bool, string, string)
}

type parser struct{}

func New() Parser {
	return &parser{}
}

func (p *parser) ParseUsernameAndCode(str string) (bool, string, string) {
	username, code := p.splitUsername(str, "-")
	if code == "" {
		return false, "", "0"
	}
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		return false, "", "0"
	}
	if codeInt < 0 || codeInt > 9999 {
		return false, "", "0"
	}
	return true, username, code
}

func (p *parser) splitUsername(str string, char string) (string, string) {
	s := strings.SplitN(str, char, 2)
	if len(s) == 2 {
		return s[0], s[1]
	}
	return "", ""
}

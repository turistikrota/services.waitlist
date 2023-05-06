package notify_helper

type Helper interface {
	BuildEmail(to string, subject string, data interface{}, recipient string, template ...string) interface{}
	BuildSMS(phone string, recipient string, text string, lang string) interface{}
	BuildTelegram(chatId string, text string, recipient string) interface{}
}

type helper struct{}

func New() Helper {
	return &helper{}
}

func (h *helper) BuildEmail(to string, subject string, data interface{}, recipient string, template ...string) interface{} {
	t := "default"
	if len(template) > 0 {
		t = template[0]
	}
	return &Mail{
		Recipient: recipient,
		To:        to,
		Subject:   subject,
		Data:      data,
		Template:  t,
	}
}

func (h *helper) BuildSMS(phone string, recipient string, text string, lang string) interface{} {
	return &SMS{
		Recipient: recipient,
		Phone:     phone,
		Text:      text,
		Lang:      lang,
	}
}

func (h *helper) BuildTelegram(chatId string, text string, recipient string) interface{} {
	return &Telegram{
		ChatId:    chatId,
		Text:      text,
		Recipient: recipient,
	}
}

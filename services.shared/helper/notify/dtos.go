package notify_helper

type Mail struct {
	Recipient string      `json:"recipient"`
	To        string      `json:"to"`
	Subject   string      `json:"subject"`
	Data      interface{} `json:"data"`
	Template  string      `json:"template"`
}

type SMS struct {
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
	Text      string `json:"text"`
	Lang      string `json:"lang"`
}

type Telegram struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	Recipient string `json:"recipient"`
}

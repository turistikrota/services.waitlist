package http

type successMessages struct {
	WaitlistJoined string
	WaitlistLeaved string
}

type errorMessages struct {
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		WaitlistJoined: "http_waitlist_joined",
		WaitlistLeaved: "http_waitlist_leaved",
	},
	Error: errorMessages{},
}

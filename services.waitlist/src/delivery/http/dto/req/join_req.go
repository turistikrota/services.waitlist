package req

import "opensource.turistikrota.com/waitlist/src/app/command"

type JoinRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (r *JoinRequest) ToCommand(lang string) command.WaitlistJoinCommand {
	return command.WaitlistJoinCommand{
		Email: r.Email,
		Lang:  lang,
	}
}

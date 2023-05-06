package req

import "opensource.turistikrota.com/waitlist/src/app/command"

type LeaveRequest struct {
	Token string `param:"token" validate:"required,uuid"`
}

func (r *LeaveRequest) ToCommand(lang string) command.WaitlistLeaveCommand {
	return command.WaitlistLeaveCommand{
		LeaveToken: r.Token,
		Lang:       lang,
	}
}

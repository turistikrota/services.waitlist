package app

import (
	"opensource.turistikrota.com/waitlist/src/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Join  command.WaitlistJoinHandler
	Leave command.WaitlistLeaveHandler
}

type Queries struct {
}

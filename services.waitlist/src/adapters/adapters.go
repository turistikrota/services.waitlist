package adapters

import (
	"opensource.turistikrota.com/waitlist/src/adapters/memory"
	"opensource.turistikrota.com/waitlist/src/adapters/mongo"
	"opensource.turistikrota.com/waitlist/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)

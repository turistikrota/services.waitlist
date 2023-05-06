package sql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Config struct {
	Address  string
	User     string
	Password string
	Database string
	Driver   string
}

func New(cnf Config) (*sqlx.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = cnf.Address
	config.User = cnf.User
	config.Passwd = cnf.Password
	config.DBName = cnf.Database
	config.ParseTime = true

	db, err := sqlx.Connect(cnf.Driver, config.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open mysql connection")
	}
	return db, nil
}

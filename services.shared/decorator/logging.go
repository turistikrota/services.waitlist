package decorator

import (
	"context"
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	WithFields(fields logrus.Fields) *logrus.Entry
	Debug(args ...interface{})
	Info(args ...interface{})
	WithError(err error) *logrus.Entry
}

func NewLogger() Logger {
	return logrus.NewEntry(logrus.StandardLogger())
}

type commandLoggingDecorator[C any, R any] struct {
	base   CommandHandler[C, R]
	logger Logger
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger Logger
}

func (d commandLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err *i18np.Error) {
	handlerType := generateActionName(cmd)

	logger := d.logger.WithFields(logrus.Fields{
		"command":      handlerType,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.WithError(err).Error("Failed to execute command")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err *i18np.Error) {
	logger := d.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.WithError(err).Error("Failed to execute query")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

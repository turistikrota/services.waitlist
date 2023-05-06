package logs

import "github.com/sirupsen/logrus"

func LogCommandExecution(command string, cmd interface{}, err error) {
	log := logrus.WithField("cmd", cmd)

	if err == nil {
		log.Infof("Command %v executed successfully", command)
		return
	}

	log.WithError(err).Errorf("Command %v failed", command)
}

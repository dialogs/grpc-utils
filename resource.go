package util

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// AcquireEventually tries to execute `acq`-function and returns its result otherwise retry within `interval`
func AcquireEventually(name string, logger *log.Entry, acq func() (interface{}, error), interval time.Duration) interface{} {
	var res interface{}
	var err error

	boundLogger := logger.WithField("resource", name)

	for {
		boundLogger.Info("Acquiring resource: pending")
		res, err = acq()
		if err == nil {
			break
		} else {
			logger.Errorf("Acquiring resource: failure [%s]. Will retry in %s", err.Error(), interval.String())
		}
		<-time.After(interval)
	}
	boundLogger.Info("Acquiring resource: ok")
	return res
}

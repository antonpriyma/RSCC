package utils

import "github.com/antonpriyma/RSCC/pkg/log"

func Must(logger log.Logger, err error, msg string) {
	if err != nil {
		logger.WithError(err).Fatal(msg)
	}
}

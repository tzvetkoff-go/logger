///usr/bin/env true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"os"
	"path"

	"github.com/tzvetkoff-go/logger"
	"github.com/tzvetkoff-go/logger/backends/syslog"
)

// revive:disable:line-length-limit

func main() {
	// As well as writing to stderr, send messages in JSON format to the system logging facility.
	backend, err := syslog.NewBackend(
		"",
		"",
		syslog.LOG_DEBUG|syslog.LOG_LOCAL7,
		path.Base(os.Args[0]),
		&logger.JSONFormatter{},
	)
	if err != nil {
		logger.Error("could not connect to syslog: %s", err)
	} else {
		logger.AddBackends(backend)
	}

	// Standard severity levels, as defined by RFC5424.
	logger.Emerg("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Alert("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Crit("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Error("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Warning("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Notice("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Info("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
	logger.Debug("This is a test log message [%d]", os.Getpid(), logger.Fields{"foo": true, "bar": "baz", "nil": nil, "empty": ""})
}

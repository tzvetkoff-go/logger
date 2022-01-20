//go:build !plan9 && !windows

// Package syslog provides a SysLog Logger backend.
package syslog

import (
	sl "log/syslog"

	"github.com/tzvetkoff-go/logger"
)

// Backend represents a SysLog Logger backend.
type Backend struct {
	Writer    *sl.Writer
	Formatter logger.Formatter
}

// NewBackend creates a new SysLog Logger backend.
func NewBackend(
	proto string,
	addr string,
	priority Priority,
	tag string,
	formatter logger.Formatter,
) (*Backend, error) {
	writer, err := sl.Dial(proto, addr, sl.Priority(priority), tag)
	if err != nil {
		return nil, err
	}

	return &Backend{
		Writer:    writer,
		Formatter: formatter,
	}, nil
}

// Write writes the log message to the backend.
func (h *Backend) Write(fields logger.Fields) bool {
	level := sl.Priority(fields["level"].(logger.Level))
	message := h.Formatter.Format(fields)

	switch level {
	case sl.LOG_EMERG:
		h.Writer.Emerg(message)
	case sl.LOG_ALERT:
		h.Writer.Alert(message)
	case sl.LOG_CRIT:
		h.Writer.Crit(message)
	case sl.LOG_ERR:
		h.Writer.Err(message)
	case sl.LOG_WARNING:
		h.Writer.Warning(message)
	case sl.LOG_NOTICE:
		h.Writer.Notice(message)
	case sl.LOG_INFO:
		h.Writer.Info(message)
	default:
		h.Writer.Debug(message)
	}

	return true
}

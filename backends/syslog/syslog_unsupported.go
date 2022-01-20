//go:build plan9 || windows

package syslog

import (
	"errors"

	"github.com/tzvetkoff-go/logger"
)

// NewBackend returns a simple error since syslog is not supported on these platforms.
func NewBackend(
	proto string,
	addr string,
	priority Priority,
	tag string,
	formatter logger.Formatter,
) (logger.Backend, error) {
	return nil, errors.New("syslog is not supported on this platform")
}

package logger

import (
	"fmt"
	"io"
)

// WriterBackend is a basic io.Writer logger backend.
type WriterBackend struct {
	Writer    io.Writer
	Formatter Formatter
}

// Write writes the log message to the backend.
func (h *WriterBackend) Write(fields Fields) bool {
	fmt.Fprintln(h.Writer, h.Formatter.Format(fields))
	return true
}

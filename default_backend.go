package logger

import (
	"io"
	"os"
)

// DefaultBackend is the default logger backend.
var DefaultBackend Backend = &WriterBackend{
	Writer:    os.Stderr,
	Formatter: DefaultFormatter,
}

// SetWriter sets the output writer of the DefaultBackend.
func SetWriter(writer io.Writer) {
	if backend, ok := DefaultBackend.(*WriterBackend); ok {
		backend.Writer = writer
	}
}

// SetFormatter sets the formatter of the DefaultBackend.
func SetFormatter(formatter Formatter) {
	if backend, ok := DefaultBackend.(*WriterBackend); ok {
		backend.Formatter = formatter
	}
}

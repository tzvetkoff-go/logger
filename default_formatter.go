package logger

import (
	"os"
)

// Stat stderr so we can figure out whether it's a terminal or not.
var stderrStat, _ = os.Stderr.Stat()

// DefaultFormatter is the default log formatter.
var DefaultFormatter Formatter = &TextFormatter{
	Timestamps:      true,
	TimestampFormat: "2006-01-02T15:04:05.000000Z: ",
	Color:           stderrStat.Mode()&os.ModeCharDevice != 0,
}

// SetTimestamps ...
func SetTimestamps(timestamps bool) {
	if formatter, ok := DefaultFormatter.(*TextFormatter); ok {
		formatter.Timestamps = timestamps
	}
}

// SetTimestampFormat ...
func SetTimestampFormat(format string) {
	if formatter, ok := DefaultFormatter.(*TextFormatter); ok {
		formatter.TimestampFormat = format
	}
}

// SetColor ...
func SetColor(color bool) {
	if formatter, ok := DefaultFormatter.(*TextFormatter); ok {
		formatter.Color = color
	}
}

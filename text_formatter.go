package logger

import (
	"fmt"
	"strings"
	"time"
)

// TextFormatter is a simple text log formatter.
type TextFormatter struct {
	Timestamps      bool
	TimestampFormat string
	Color           bool
}

// Format returns log fields formatted as a string.
func (f *TextFormatter) Format(fields Fields) string {
	level := fields["level"].(Level)
	delete(fields, "level")

	t := fields["time"].(time.Time)
	delete(fields, "time")

	msg := fields["msg"].(string)
	delete(fields, "msg")

	color := ""
	reset := ""
	if f.Color {
		switch level {
		case LOG_EMERG:
			fallthrough
		case LOG_ALERT:
			fallthrough
		case LOG_CRIT:
			fallthrough
		case LOG_ERROR:
			color = "\x1b[1;31m" // bright-red
			reset = "\x1b[0;0m"
		case LOG_WARNING:
			color = "\x1b[1;33m" // bright-yellow
			reset = "\x1b[0;0m"
		case LOG_NOTICE:
			color = "\x1b[1;37m" // bright-white
			reset = "\x1b[0;0m"
		case LOG_INFO:
			color = "\x1b[0;0m" // normal
			reset = "\x1b[0;0m"
		case LOG_DEBUG:
			color = "\x1b[1;30m" // gray
			reset = "\x1b[0;0m"
		}
	}

	builder := &strings.Builder{}

	if f.Timestamps {
		builder.WriteString(t.UTC().Format(f.TimestampFormat))
	}

	fmt.Fprintf(builder, "%s%s%s: %-48s", color, LookupLevel(level)[0:4], reset, msg)

	for _, k := range fields.Keys() {
		v := fields[k]

		if vv, ok := v.(string); ok {
			fmt.Fprintf(builder, " %s=%q", k, vv)
		} else {
			fmt.Fprint(builder, " ", k, "=", v)
		}
	}

	return strings.TrimSpace(builder.String())
}

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
	t := fields["time"].(time.Time)
	delete(fields, "time")

	level := fields["level"].(Level)
	delete(fields, "level")

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

	if f.Color {
		builder.WriteString(color)
	}

	if f.Timestamps {
		fmt.Fprintf(builder, "time=%q ", t.UTC().Format(f.TimestampFormat))
	}

	fmt.Fprintf(builder, "level=%q", LookupLevel(level)[0:4])

	if msg != "" {
		fmt.Fprintf(builder, " msg=%q", msg)
	}

	for _, k := range fields.Keys() {
		v := fields[k]

		if vv, ok := v.(string); ok {
			fmt.Fprintf(builder, " %s=%q", k, vv)
		} else {
			fmt.Fprint(builder, " ", k, "=", v)
		}
	}

	if f.Color {
		builder.WriteString(reset)
	}

	return builder.String()
}

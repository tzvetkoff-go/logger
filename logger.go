package logger

import (
	"fmt"
	"time"
)

// Logger is the logger itself.
type Logger struct {
	Level    Level
	Fields   Fields
	Backends []Backend
}

// Log ...
func (l *Logger) Log(level Level, format string, v ...interface{}) {
	if l.Level >= level {
		fields := l.Fields.Copy()
		if len(v) > 0 {
			if f, ok := v[len(v)-1].(Fields); ok {
				fields.Update(f)
				v = v[0 : len(v)-1]
			}
		}

		fields["level"] = level
		fields["time"] = time.Now()
		fields["msg"] = fmt.Sprintf(format, v...)

		for _, backend := range l.Backends {
			if !backend.Write(fields.Copy()) {
				break
			}
		}
	}
}

// Emerg ...
func (l *Logger) Emerg(format string, v ...interface{}) {
	l.Log(LOG_EMERG, format, v...)
}

// Alert ...
func (l *Logger) Alert(format string, v ...interface{}) {
	l.Log(LOG_ALERT, format, v...)
}

// Crit ...
func (l *Logger) Crit(format string, v ...interface{}) {
	l.Log(LOG_CRIT, format, v...)
}

// Err ...
func (l *Logger) Error(format string, v ...interface{}) {
	l.Log(LOG_ERROR, format, v...)
}

// Warning ...
func (l *Logger) Warning(format string, v ...interface{}) {
	l.Log(LOG_WARNING, format, v...)
}

// Notice ...
func (l *Logger) Notice(format string, v ...interface{}) {
	l.Log(LOG_NOTICE, format, v...)
}

// Info ...
func (l *Logger) Info(format string, v ...interface{}) {
	l.Log(LOG_INFO, format, v...)
}

// Debug ...
func (l *Logger) Debug(format string, v ...interface{}) {
	l.Log(LOG_DEBUG, format, v...)
}

package logger

import (
	"strings"
)

// Level represents a log level.
type Level int

// Levels as defined by RFC5424.
const (
	// revive:disable:var-naming
	LOG_EMERG   = Level(0)
	LOG_ALERT   = Level(1)
	LOG_CRIT    = Level(2)
	LOG_ERROR   = Level(3)
	LOG_WARNING = Level(4)
	LOG_NOTICE  = Level(5)
	LOG_INFO    = Level(6)
	LOG_DEBUG   = Level(7)
	// revive:enable:var-naming
)

// StringToLevel maps common string names to log levels.
var StringToLevel = map[string]Level{
	"LOG_EMERG":   LOG_EMERG,
	"EMERG":       LOG_EMERG,
	"emerg":       LOG_EMERG,
	"0":           LOG_EMERG,
	"LOG_ALERT":   LOG_ALERT,
	"ALERT":       LOG_ALERT,
	"alert":       LOG_ALERT,
	"1":           LOG_ALERT,
	"LOG_CRIT":    LOG_CRIT,
	"CRIT":        LOG_CRIT,
	"crit":        LOG_CRIT,
	"2":           LOG_CRIT,
	"LOG_ERROR":   LOG_ERROR,
	"ERROR":       LOG_ERROR,
	"error":       LOG_ERROR,
	"3":           LOG_ERROR,
	"LOG_WARNING": LOG_WARNING,
	"WARNING":     LOG_WARNING,
	"warning":     LOG_WARNING,
	"4":           LOG_WARNING,
	"LOG_NOTICE":  LOG_NOTICE,
	"NOTICE":      LOG_NOTICE,
	"notice":      LOG_NOTICE,
	"5":           LOG_NOTICE,
	"LOG_INFO":    LOG_INFO,
	"INFO":        LOG_INFO,
	"info":        LOG_INFO,
	"6":           LOG_INFO,
	"LOG_DEBUG":   LOG_DEBUG,
	"DEBUG":       LOG_DEBUG,
	"debug":       LOG_DEBUG,
	"7":           LOG_DEBUG,
}

// LevelToString maps levels to string names.
var LevelToString = map[Level]string{
	LOG_EMERG:   "EMERG",
	LOG_ALERT:   "ALERT",
	LOG_CRIT:    "CRIT",
	LOG_ERROR:   "ERROR",
	LOG_WARNING: "WARNING",
	LOG_NOTICE:  "NOTICE",
	LOG_INFO:    "INFO",
	LOG_DEBUG:   "DEBUG",
}

// LookupString returns a log level from a string.
func LookupString(s string) Level {
	result := Level(0)

	names := strings.FieldsFunc(s, func(r rune) bool {
		return r == ',' || r == '|' || r == ' ' || r == '\t'
	})
	for _, name := range names {
		if v, ok := StringToLevel[name]; ok {
			result |= v
		}
	}

	return result
}

// LookupLevel returns a string from a log level.
func LookupLevel(c Level) string {
	if s, ok := LevelToString[c&0x07]; ok {
		return s
	}

	return ""
}

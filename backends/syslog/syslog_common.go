package syslog

import "strings"

// Priority describes SysLog severities and logging facilities.
type Priority int

// Priority constants.
const ( // revive:disable:var-naming
	// Severities.
	LOG_EMERG   Priority = 0
	LOG_ALERT   Priority = 1
	LOG_CRIT    Priority = 2
	LOG_ERR     Priority = 3
	LOG_WARNING Priority = 4
	LOG_NOTICE  Priority = 5
	LOG_INFO    Priority = 6
	LOG_DEBUG   Priority = 7

	// Logging facilities.
	LOG_KERN     Priority = 0
	LOG_USER     Priority = 8
	LOG_MAIL     Priority = 16
	LOG_DAEMON   Priority = 24
	LOG_AUTH     Priority = 32
	LOG_SYSLOG   Priority = 40
	LOG_LPR      Priority = 48
	LOG_NEWS     Priority = 56
	LOG_UUCP     Priority = 64
	LOG_CRON     Priority = 72
	LOG_AUTHPRIV Priority = 80
	LOG_FTP      Priority = 88

	LOG_LOCAL0 Priority = 128
	LOG_LOCAL1 Priority = 136
	LOG_LOCAL2 Priority = 144
	LOG_LOCAL3 Priority = 152
	LOG_LOCAL4 Priority = 160
	LOG_LOCAL5 Priority = 168
	LOG_LOCAL6 Priority = 176
	LOG_LOCAL7 Priority = 184
) // revive:enable:var-naming

// StringToPriority ...
var StringToPriority = map[string]Priority{
	// Severities ...
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
	"LOG_ERR":     LOG_ERR,
	"ERR":         LOG_ERR,
	"err":         LOG_ERR,
	"3":           LOG_ERR,
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
	// Local facilities ...
	"LOG_LOCAL0": LOG_LOCAL0,
	"LOCAL0":     LOG_LOCAL0,
	"local0":     LOG_LOCAL0,
	"LOG_LOCAL1": LOG_LOCAL1,
	"LOCAL1":     LOG_LOCAL1,
	"local1":     LOG_LOCAL1,
	"LOG_LOCAL2": LOG_LOCAL2,
	"LOCAL2":     LOG_LOCAL2,
	"local2":     LOG_LOCAL2,
	"LOG_LOCAL3": LOG_LOCAL3,
	"LOCAL3":     LOG_LOCAL3,
	"local3":     LOG_LOCAL3,
	"LOG_LOCAL4": LOG_LOCAL4,
	"LOCAL4":     LOG_LOCAL4,
	"local4":     LOG_LOCAL4,
	"LOG_LOCAL5": LOG_LOCAL5,
	"LOCAL5":     LOG_LOCAL5,
	"local5":     LOG_LOCAL5,
	"LOG_LOCAL6": LOG_LOCAL6,
	"LOCAL6":     LOG_LOCAL6,
	"local6":     LOG_LOCAL6,
	"LOG_LOCAL7": LOG_LOCAL7,
	"LOCAL7":     LOG_LOCAL7,
	"local7":     LOG_LOCAL7,
}

// LookupString ...
func LookupString(s string) Priority {
	result := Priority(0)

	names := strings.FieldsFunc(s, func(r rune) bool {
		return r == ',' || r == '|' || r == ' ' || r == '\t'
	})
	for _, name := range names {
		if v, ok := StringToPriority[name]; ok {
			result |= v
		}
	}

	return result
}

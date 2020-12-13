package logger

// DefaultLogger is the default logger instance.
var DefaultLogger = &Logger{
	Level:  LOG_DEBUG,
	Fields: Fields{},
	Backends: []Backend{
		DefaultBackend,
	},
}

// Log ...
func Log(level Level, format string, v ...interface{}) {
	DefaultLogger.Log(level, format, v...)
}

// Emerg ...
func Emerg(format string, v ...interface{}) {
	DefaultLogger.Emerg(format, v...)
}

// Alert ...
func Alert(format string, v ...interface{}) {
	DefaultLogger.Alert(format, v...)
}

// Crit ...
func Crit(format string, v ...interface{}) {
	DefaultLogger.Crit(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	DefaultLogger.Error(format, v...)
}

// Warning ...
func Warning(format string, v ...interface{}) {
	DefaultLogger.Warning(format, v...)
}

// Notice ...
func Notice(format string, v ...interface{}) {
	DefaultLogger.Notice(format, v...)
}

// Info ...
func Info(format string, v ...interface{}) {
	DefaultLogger.Info(format, v...)
}

// Debug ...
func Debug(format string, v ...interface{}) {
	DefaultLogger.Debug(format, v...)
}

// GetLevel ...
func GetLevel() Level {
	return DefaultLogger.Level
}

// SetLevel ...
func SetLevel(level Level) {
	DefaultLogger.Level = level
}

// GetFields ...
func GetFields() Fields {
	return DefaultLogger.Fields
}

// SetFields ...
func SetFields(fields Fields) {
	DefaultLogger.Fields = fields
}

// GetBackends ...
func GetBackends() []Backend {
	return DefaultLogger.Backends
}

// AddBackends ...
func AddBackends(backends ...Backend) {
	DefaultLogger.Backends = append(DefaultLogger.Backends, backends...)
}

// ClearBackends ...
func ClearBackends() {
	DefaultLogger.Backends = []Backend{}
}

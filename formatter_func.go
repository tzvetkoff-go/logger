package logger

// FormatterFunc is a function wrapper of Formatter.
type FormatterFunc func(Fields) string

// Format returns log fields formatted as a string.
func (f FormatterFunc) Format(fields Fields) string {
	return f(fields)
}

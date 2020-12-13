package logger

// Formatter describes a log formatter.
type Formatter interface {
	Format(Fields) string
}

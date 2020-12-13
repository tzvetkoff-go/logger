package logger

// Backend describes a logger backend.
type Backend interface {
	Write(Fields) bool
}

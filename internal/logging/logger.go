package logging

// Logger levels information
type Logger interface {
	Fatal(format string, content ...interface{})
	Error(format string, content ...interface{})
	Warn(format string, content ...interface{})
	Info(format string, content ...interface{})
	Debug(format string, content ...interface{})
}

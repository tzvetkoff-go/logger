// Package log contains a compatibility layer for the system "log" package.
package log

import (
	"fmt"
	"io"
	"os"

	"github.com/tzvetkoff-go/logger"
)

// Print ...
func Print(v ...interface{}) {
	logger.Info(fmt.Sprint(v...))
}

// Println ...
func Println(v ...interface{}) {
	logger.Info(fmt.Sprintln(v...))
}

// Printf ...
func Printf(format string, v ...interface{}) {
	logger.Info(format, v...)
}

// Fatal ...
func Fatal(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalln ...
func Fatalln(v ...interface{}) {
	logger.Error(fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf ...
func Fatalf(format string, v ...interface{}) {
	logger.Error(format, v...)
	os.Exit(1)
}

// Panic ...
func Panic(v ...interface{}) {
	message := fmt.Sprint(v...)
	logger.Error(message)
	panic(message)
}

// Panicln ...
func Panicln(v ...interface{}) {
	message := fmt.Sprintln(v...)
	logger.Error(message)
	panic(message)
}

// Panicf ...
func Panicf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	logger.Error(message)
	panic(message)
}

// SetOutput ...
func SetOutput(writer io.Writer) {
	logger.SetWriter(writer)
}

package cli

import (
	"fmt"
	"io"
	"os"
)

// Print levels go in order: Debug, Info, Warn, Error, Fatal
const (
	LevelDebug = iota // i
	LevelInfo         // +
	LevelWarn         // *
	LevelError        // -
	LevelFatal        // !
)

var printLevel = LevelInfo
var outWriter io.Writer = os.Stdout
var errWriter io.Writer = os.Stderr

// SetPrintLevel allows you to set the level to print, by default LevelInfo is set
func SetPrintLevel(level int) {
	printLevel = level
}

// SetOutputWriter allows you to set the output file for debug, info, and warn messges
func SetOutputWriter(w io.Writer) {
	outWriter = w
}

// SetErrorWriter allows you to set the output writer for error and fatal messages
func SetErrorWriter(w io.Writer) {
	errWriter = w
}

// Debug prints a formatted debug level message with a newline appended
func Debug(format string, a ...interface{}) {
	writeMessage(LevelDebug, outWriter, fmt.Sprintln("[i]", fmt.Sprintf(format, a...)))
}

// Info prints a formatted info level message with a newline appended
func Info(format string, a ...interface{}) {
	writeMessage(LevelInfo, outWriter, fmt.Sprintln("[+]", fmt.Sprintf(format, a...)))
}

// Warn prints a formatted warning level message with a newline appended
func Warn(format string, a ...interface{}) {
	writeMessage(LevelWarn, outWriter, fmt.Sprintln("[*]", fmt.Sprintf(format, a...)))
}

// Error prints a formatted error level message with a newline appended
func Error(format string, a ...interface{}) {
	writeMessage(LevelError, errWriter, fmt.Sprintln("[-]", fmt.Sprintf(format, a...)))
}

// Fatal prints a formatted fatal level message with a newline appended and calls os.Exit(1)
func Fatal(format string, a ...interface{}) {
	writeMessage(LevelFatal, errWriter, fmt.Sprintln("[!]", fmt.Sprintf(format, a...)))
	os.Exit(1)
}

// Debugln prints a debug level message with a newline appended
func Debugln(a ...interface{}) {
	writeMessage(LevelDebug, outWriter, fmt.Sprintln("[i]", fmt.Sprint(a...)))
}

// Infoln prints an info level message with a newline appended
func Infoln(a ...interface{}) {
	writeMessage(LevelInfo, outWriter, fmt.Sprintln("[+]", fmt.Sprint(a...)))
}

// Warnln prints a warning level message with a newline appended
func Warnln(a ...interface{}) {
	writeMessage(LevelWarn, outWriter, fmt.Sprintln("[*]", fmt.Sprint(a...)))
}

// Errorln prints an error level message with a newline appended
func Errorln(a ...interface{}) {
	writeMessage(LevelError, outWriter, fmt.Sprintln("[-]", fmt.Sprint(a...)))
}

// Fatalln prints a fatal level message with a newline appended and calls os.Exit(1)
func Fatalln(a ...interface{}) {
	writeMessage(LevelFatal, outWriter, fmt.Sprintln("[!]", fmt.Sprint(a...)))
	os.Exit(1)
}

// WriteBanner is a special function to be used only when printing the program banner
func WriteBanner(a string) {
	writeMessage(LevelInfo, outWriter, fmt.Sprintln(a))
}

// WriteResults is meant to write results of tool actions
func WriteResults(a ...interface{}) {
	writeMessage(6, outWriter, fmt.Sprintln(a...))
}

func writeMessage(level int, writer io.Writer, message string) {
	if level < printLevel {
		return
	}
	fmt.Fprint(writer, message)
}

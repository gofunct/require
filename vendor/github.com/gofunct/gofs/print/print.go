package print

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
	"go.uber.org/zap"
	"io"
	"sync"
	"github.com/pkg/errors"
)

var cyan func(string) string
var red func(string) string
var yellow func(string) string
var redInverse func(string) string
var gray func(string) string
var magenta func(string) string

var colorfulMap = map[string]int{}
var colorfulMutex = &sync.Mutex{}
var colorfulFormats = []func(string) string{
	ansi.ColorFunc("+h"),
	ansi.ColorFunc("green"),
	ansi.ColorFunc("yellow"),
	ansi.ColorFunc("magenta"),
	ansi.ColorFunc("green+h"),
	ansi.ColorFunc("yellow+h"),
	ansi.ColorFunc("magenta+h"),
}

// LogWriter is the writer to which the logs are written
type syncer struct {
	io.Writer
}

func (l *syncer) Sync() error {
	return nil
}

var LogWriter *syncer

func init() {
	ansi.DisableColors(false)
	cyan = ansi.ColorFunc("cyan")
	red = ansi.ColorFunc("red+b")
	yellow = ansi.ColorFunc("yellow+b")
	redInverse = ansi.ColorFunc("white:red")
	gray = ansi.ColorFunc("black+h")
	magenta = ansi.ColorFunc("magenta+h")
	LogWriter = &syncer{
		colorable.NewColorableStdout(),
	}
	lger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(lger)
	zap.ErrorOutput(LogWriter)
}

// Debug writes a debug statement to stdout.
func Debug(group string, format string, any ...interface{}) {
	_, _ =  fmt.Fprint(LogWriter, gray(group)+" ")
	_, _ = fmt.Fprintf(LogWriter, gray(format))
}

// Info writes an info statement to stdout.
func Info(group string, format string, any ...interface{}) {
	_, _ = fmt.Fprint(LogWriter, cyan(group)+" ")
	_, _ = fmt.Fprintf(LogWriter, format, any...)
}

// InfoColorful writes an info statement to stdout changing colors
// on succession.
func InfoColorful(group string, format string, any ...interface{}) {
	colorfulMutex.Lock()
	colorfulMap[group]++
	colorFn := colorfulFormats[colorfulMap[group]%len(colorfulFormats)]
	colorfulMutex.Unlock()
	_, _ = fmt.Fprint(LogWriter, cyan(group)+" ")
	s := colorFn(fmt.Sprintf(format, any...))

	_, _ =  fmt.Fprint(LogWriter, s)
}

// Error writes an error statement to stdout.
func Error(group string, format string, any ...interface{}) error {
	_, _ =fmt.Fprintf(LogWriter, red(group)+" ")
	_ , _ = fmt.Fprintf(LogWriter, red(format), any...)

	return fmt.Errorf(format, any...)
}

// Panic writes an error statement to stdout.
func Panic(group string, format string, any ...interface{}) {
	_, _ = fmt.Fprintf(LogWriter, redInverse(group)+" ")

	_, _ = fmt.Fprintf(LogWriter, redInverse(format), any...)

	panic("")
}

// Error writes an error statement to stdout.
func WithStack(e error) error {
	return errors.WithStack(e)
}

// Error writes an error statement to stdout.
func Hint(e error, hint string) error {
	return errors.Wrap(e, hint)
}
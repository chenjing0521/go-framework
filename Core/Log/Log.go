package Log

import (
	"fmt"
	"io"
	"os"
	"time"
)

// recode log
type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	WriteFile(fileName, content string)
}

//logger struct
type MyLogger struct {
	out io.Writer
}

//implement Logger function
func (l *MyLogger) Print(args ...interface{}) {
	if l.out == nil {
		l.out = os.Stdout
	}
	fmt.Print(time.Now().Format("2006-01-02 15:04:05 - "))
	fmt.Fprintln(l.out, args...)
}

//implement Logger function
func (l *MyLogger) Printf(format string, args ...interface{}) {
	l.Print(fmt.Sprintf(format, args...))
}

func (l *MyLogger) WriteFile(fileName, content string) {
	//determine if the file exists

}

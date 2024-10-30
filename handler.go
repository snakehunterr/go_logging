package logging

import (
	"fmt"
	"io"
	"os"
)

type Handler interface {
	Handle(*LoggerMessage)
}

type WriterHandler struct {
	Level     LogLevel
	Writer    io.Writer
	Formatter Formatter
}

func (h WriterHandler) Handle(msg *LoggerMessage) {
	if h.Level > msg.Level {
		return
	}
	if _, err := fmt.Fprintln(h.Writer, h.Formatter.Format(msg)); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}

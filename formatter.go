package logging

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Formatter interface {
	Format(*LoggerMessage) string
}

type BasicFormatter struct {
	MessageFormat string
	TimeFormat    string
}

func (f BasicFormatter) Format(msg *LoggerMessage) string {
	s := f.MessageFormat

	s = strings.ReplaceAll(s, "%(level)s", msg.LevelString)
	s = strings.ReplaceAll(s, "%(name)s", msg.LoggerName)
	s = strings.ReplaceAll(s, "%(time)s", msg.Time.Format(f.TimeFormat))
	s = strings.ReplaceAll(s, "%(msg)s", fmt.Sprint(msg.Message...))

	return s
}

type ColoredFormatter struct {
	MessageFormat string
	TimeFormat    string
	Colors        map[LogLevel]lipgloss.Style
}

func (f ColoredFormatter) Format(msg *LoggerMessage) string {
	s := f.MessageFormat
	style := f.Colors[msg.Level]

	s = strings.ReplaceAll(s, "%(level)s", style.Render(msg.LevelString))
	s = strings.ReplaceAll(s, "%(name)s", MessageLoggerNameStyle.Render(msg.LoggerName))
	s = strings.ReplaceAll(s, "%(time)s", MessageTimeStyle.Render(msg.Time.Format(f.TimeFormat)))
	s = strings.ReplaceAll(s, "%(msg)s", fmt.Sprint(msg.Message...))

	return s
}

var (
	DefaultColoredFormatter = ColoredFormatter{
		MessageFormat: "%(level)s %(name)s :: %(time)s :: %(msg)s",
		TimeFormat:    "2006 January 02 - 15:04:05",
		Colors:        DefaultColors,
	}
)

var (
	DefaultFormatter = BasicFormatter{
		MessageFormat: "%(level)s\t%(name)s :: %(time)s :: %(msg)s",
		TimeFormat:    "2006 January 02 - 15:04:05",
	}
)

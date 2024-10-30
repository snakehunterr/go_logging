package logging

import "time"

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelLog
	LevelWarning
	LevelError
)

var (
	LevelMap = map[LogLevel]string{
		LevelDebug:   "DEBUG",
		LevelLog:     "LOG",
		LevelWarning: "WARNING",
		LevelError:   "ERROR",
	}
)

type LoggerMessage struct {
	Level       LogLevel
	LevelString string
	LoggerName  string
	Time        time.Time
	Message     []any
}

type Logger struct {
	Level    LogLevel
	Name     string
	Handlers []Handler
}

func (l Logger) log(level LogLevel, v ...any) {
	var msg = &LoggerMessage{
		Level:       level,
		LevelString: LevelMap[level],
		LoggerName:  l.Name,
		Time:        time.Now(),
		Message:     v,
	}
	for _, h := range l.Handlers {
		h.Handle(msg)
	}
}

func (l Logger) Error(v ...any) {
	l.log(LevelError, v...)
}

func (l Logger) Warning(v ...any) {
	if l.Level > LevelWarning {
		return
	}

	l.log(LevelWarning, v...)
}

func (l Logger) Log(v ...any) {
	if l.Level > LevelLog {
		return
	}

	l.log(LevelLog, v...)
}

func (l Logger) Debug(v ...any) {
	if l.Level > LevelDebug {
		return
	}

	l.log(LevelDebug, v...)
}

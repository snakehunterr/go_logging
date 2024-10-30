package logging

import "github.com/charmbracelet/lipgloss"

var (
	MessageTimeColor       = "#ff9a01"
	MessageLoggerNameColor = "#4fa4ff"

	MessageTimeStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color(MessageTimeColor))
	MessageLoggerNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(MessageLoggerNameColor))

	BlackColor               = "#131313"
	MessageLevelDebugColor   = "#87f5ff"
	MessageLevelLogColor     = "#49ff49"
	MessageLevelWarningColor = "#fdbe4c"
	MessageLevelErrorColor   = "#ff4f4f"

	MessageLevelWidth = 9

	MessageLevelDebugStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(MessageLevelDebugColor)).
				Foreground(lipgloss.Color(BlackColor)).
				Width(MessageLevelWidth).
				Align(lipgloss.Center)

	MessageLevelLogStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(MessageLevelLogColor)).
				Foreground(lipgloss.Color(BlackColor)).
				Width(MessageLevelWidth).
				Align(lipgloss.Center)

	MessageLevelWarningStyle = lipgloss.NewStyle().
					Background(lipgloss.Color(MessageLevelWarningColor)).
					Foreground(lipgloss.Color(BlackColor)).
					Width(MessageLevelWidth).
					Align(lipgloss.Center)

	MessageLevelErrorStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(MessageLevelErrorColor)).
				Foreground(lipgloss.Color(BlackColor)).
				Width(MessageLevelWidth).
				Align(lipgloss.Center)

	DefaultColors = map[LogLevel]lipgloss.Style{
		LevelDebug:   MessageLevelDebugStyle,
		LevelLog:     MessageLevelLogStyle,
		LevelWarning: MessageLevelWarningStyle,
		LevelError:   MessageLevelErrorStyle,
	}
)

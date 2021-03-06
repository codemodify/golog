package persisters

import (
	"fmt"

	gologC "github.com/brightappsllc/golog/contracts"
	coloredLogs "github.com/logrusorgru/aurora"
)

type consoleLogger struct {
	logUntil int
}

// NewConsoleLogger -
func NewConsoleLogger(logUntil int) gologC.Logger {
	return &consoleLogger{
		logUntil: logUntil,
	}
}

func (thisRef consoleLogger) Log(logEntry gologC.LogEntry) {
	if logEntry.LogType == gologC.TypeDisable {
		return
	}

	if logEntry.LogType > thisRef.logUntil &&
		logEntry.LogType != gologC.TypeTrace {
		return
	}

	if logEntry.LogType < gologC.TypeWarning {
		fmt.Println(coloredLogs.Red(logEntry.Message))
	} else if logEntry.LogType == gologC.TypeWarning {
		fmt.Println(coloredLogs.Magenta(logEntry.Message))
	} else if logEntry.LogType == gologC.TypeInfo {
		fmt.Println(logEntry.Message)
	} else if logEntry.LogType == gologC.TypeDebug {
		fmt.Println(coloredLogs.Green(logEntry.Message))
	}
}

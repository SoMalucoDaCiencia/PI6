package src

import (
	"fmt"
	"time"
)

type MsgKind string

type String []byte

const (
	LogOk   MsgKind = "SUCCESS"
	LogInfo MsgKind = "INFO"
	LogWarn MsgKind = "WARN"
	LogErr  MsgKind = "ERR"

	// Colors
	dGray  string = "\033[90m"
	lGray  string = "\033[37m"
	blue   string = "\033[34m"
	yellow string = "\033[33m"
	green  string = "\033[32m"
	red    string = "\033[31m"
	bold   string = "\033[1m"
	reset  string = "\033[0m"
)

func WriteLog(kind MsgKind, msg string, origin string) {
	if len(origin) == 0 {
		origin = "internal"
	}
	color := green
	switch kind {
	case LogOk:
		color = green
		break
	case LogInfo:
		color = blue
		break
	case LogWarn:
		color = yellow
		break
	case LogErr:
		color = red
		break
	default:
		color = red
		break
	}
	fmt.Printf("%s|%s|%s - %s%s%s%s :: %s - %s%s%s\n", color, kind, reset, bold, lGray, origin, reset, msg, bold, time.Now().Format("02-Jan-2006 03:04 PM"), reset)
}

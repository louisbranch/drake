package presenter

import (
	"math"
	"time"

	"golang.org/x/text/message"

	"github.com/louisbranch/drake"
)

type Session struct {
	drake.Session
	ElapsedTime string
}

func SessionsList(sessions []drake.Session, printer *message.Printer) []Session {
	list := []Session{}

	for _, session := range sessions {
		var elapsed string

		mins := time.Since(session.CreatedAt).Minutes()
		hours := time.Since(session.CreatedAt).Hours()

		switch {
		case mins < 1:
			elapsed = printer.Sprintf("Less than one min ago")
		case mins < 90:
			min := int64(math.Ceil(mins))
			elapsed = printer.Sprintf("%d mins ago", min)
		case hours < 24:
			hr := int64(math.Ceil(hours))
			elapsed = printer.Sprintf("%d hours ago", hr)
		default:
			days := int64(math.Ceil(hours / 24))
			elapsed = printer.Sprintf("%d days ago", days)
		}

		list = append(list, Session{session, elapsed})
	}

	return list
}

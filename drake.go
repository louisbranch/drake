package drake

import (
	"math"
	"time"
)

type Session struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type Survey struct {
	ID          string
	SessionID   string
	AccessToken string

	R  float64
	Fp float64
	Ne float64
	Fl float64
	Fi float64
	Fc float64
	L  float64
	N  float64

	PresurveyFamiliarity bool
	PresurveyAssessment  float64
	PostsurveyDifference int64
	PostsurveyLearnGain  int64
	PostsurveyReason     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Session) MinutesSince() int64 {
	mins := time.Since(s.CreatedAt).Minutes()
	return int64(math.Floor(mins))
}

type Database interface {
	CreateSession(*Session) error
	FindSessions() ([]Session, error)
	FindSession(string) (Session, error)
}

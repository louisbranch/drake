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

func (s Session) MinutesSince() int64 {
	mins := time.Since(s.CreatedAt).Minutes()
	return int64(math.Floor(mins))
}

type Survey struct {
	ID          string
	SessionID   string
	AccessToken string

	R  *float64
	Fp *float64
	Ne *float64
	Fl *float64
	Fi *float64
	Fc *float64
	L  *float64
	N  *float64

	PresurveyFamiliarity *bool
	PresurveyAssessment  *int64
	PostsurveyDifference *int64
	PostsurveyLearnGain  *int64
	PostsurveyReason     *string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Survey) Result() float64 {
	// TODO: round?
	return *s.R * *s.Fp * *s.Ne * *s.Fl * *s.Fi * *s.Fc * *s.Fc
}

type Database interface {
	CreateSession(*Session) error
	FindSessions() ([]Session, error)
	FindSession(string) (Session, error)

	CreateSurvey(*Survey) error
	UpdateSurvey(*Survey) error
	FindSurvey(string, string) (Survey, error)
}

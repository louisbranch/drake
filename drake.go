package drake

import (
	"time"
)

type Session struct {
	ID           string
	Name         string
	Participants int64
	CreatedAt    time.Time
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

	PresurveyAssessment *float64
	PostsurveyLearnGain *bool
	PostsurveyReason    *string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Survey) Result() {
	n := *s.R * *s.Fp * *s.Ne * *s.Fl * *s.Fi * *s.Fc * *s.Fc
	s.N = &n
}

func (s Survey) NextQuestion() string {
	switch {
	case s.PresurveyAssessment == nil:
		return "Assessment"
	case s.R == nil:
		return "R"
	case s.Fp == nil:
		return "Fp"
	case s.Ne == nil:
		return "Ne"
	case s.Fl == nil:
		return "Fl"
	case s.Fi == nil:
		return "Fi"
	case s.Fc == nil:
		return "Fc"
	case s.L == nil:
		return "L"
	case s.PostsurveyLearnGain == nil:
		return "LearnGain"
	default:
		return ""
	}
}

type Database interface {
	CreateSession(*Session) error
	FindSessions() ([]Session, error)
	FindSession(string) (Session, error)

	CreateSurvey(*Survey) error
	UpdateSurvey(*Survey) error
	FindSurvey(string, string) (Survey, error)
	FindSurveys(string) ([]Survey, error)
}

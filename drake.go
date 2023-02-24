package drake

import (
	"math"
	"time"
)

type Session struct {
	ID           string
	Name         string
	Participants int64
	CreatedAt    time.Time
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

	PresurveyAssessment *float64
	PostsurveyLearnGain *bool
	PostsurveyReason    *string

	CreatedAt time.Time
	UpdatedAt time.Time
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

func (s *Survey) Result() {
	n := *s.R * *s.Fp * *s.Ne * *s.Fl * *s.Fi * *s.Fc * *s.Fc
	s.N = &n
}

func (s Survey) Difference() int64 {
	n := math.Log10(*s.N)
	a := math.Log10(*s.PresurveyAssessment)
	diff := math.Floor(math.Abs(n - a))
	return int64(diff)
}

type Result struct {
	Surveys []Survey
}

func (r Result) Participants() int {
	return len(r.Surveys)
}

func (r Result) Buckets() []string {
	supers := []rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}
	buckets := make([]string, 14)

	for i := 0; i < 14; i++ {
		val := "10"
		if i > 10 {
			val += string(supers[1]) + string(supers[i%10])
		} else {
			val += string(supers[i%10])
		}
		buckets[i] = val
	}

	return buckets
}

func (r Result) PresurveyData() []int {
	data := make([]int, len(r.Buckets()))

	for _, s := range r.Surveys {
		if s.PresurveyAssessment == nil {
			continue
		}
		n := int(math.Log10(*s.PresurveyAssessment))
		data[n] += 1
	}

	return data
}

func (r Result) PostsurveyData() []int {
	data := make([]int, len(r.Buckets()))

	for _, s := range r.Surveys {
		if s.N == nil {
			continue
		}
		n := int(math.Log10(*s.N))
		data[n] += 1
	}

	return data
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

package presenter

import (
	"math"

	"github.com/louisbranch/drake"
)

type Survey struct {
	drake.Survey
}

func (s Survey) Difference() int64 {
	if s.N == nil || s.PresurveyAssessment == nil {
		return 0
	}

	n := math.Log10(*s.N)
	if n < 1 {
		n = 0
	}
	a := math.Log10(*s.PresurveyAssessment)
	diff := math.Floor(math.Abs(n - a))
	return int64(diff)
}

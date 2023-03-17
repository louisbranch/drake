package presenter

import (
	"math"

	"github.com/louisbranch/drake"
)

type Result struct {
	Surveys []drake.Survey
}

func (r Result) Participants() int {
	return len(r.Surveys)
}

func (r Result) Buckets() []string {
	max := 15
	supers := []rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}
	buckets := make([]string, max)

	for i := 0; i < max; i++ {
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
		if n < 1 {
			n = 0
		}
		data[n] += 1
	}

	return data
}

func (r Result) PostsurveyData() []int {
	max := len(r.Buckets())
	data := make([]int, max)

	for _, s := range r.Surveys {
		if s.N == nil {
			continue
		}
		n := int(math.Log10(*s.N))
		if n < 1 {
			n = 0
		} else if n > max-1 {
			n = max - 1
		}
		data[n] += 1
	}

	return data
}

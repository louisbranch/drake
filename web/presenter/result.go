package presenter

import (
	"math"

	"github.com/louisbranch/drake"
	"golang.org/x/text/message"
)

type Result struct {
	Surveys []drake.Survey
	*message.Printer
}

func (r Result) Participants() int {
	return len(r.Surveys)
}

func (r Result) DataLabels() []string {
	return []string{
		r.Printer.Sprintf("Only us"),
		r.Printer.Sprintf("Dozens"),
		r.Printer.Sprintf("Hundreds"),
		r.Printer.Sprintf("Thousands"),
		r.Printer.Sprintf("Millions"),
	}
}

func (r Result) PresurveyData() []int {
	data := make([]int, 7)

	for _, s := range r.Surveys {
		if s.PresurveyAssessment == nil {
			continue
		}
		n := int(math.Log10(*s.PresurveyAssessment))
		switch {
		case n < 1:
			data[0] += 1 // 0
		case n < 3:
			data[n] += 1 // 1 - 999
		case n < 6:
			data[3] += 1 // 1,000 - 999,999
		default:
			data[4] += 1 // 1,000,000+
		}
	}

	return data
}

func (r Result) PostsurveyData() []int {
	data := make([]int, 7)

	for _, s := range r.Surveys {
		if s.N == nil {
			continue
		}
		n := int(math.Log10(*s.N))
		switch {
		case n < 1:
			data[0] += 1 // 0
		case n < 3:
			data[n] += 1 // 1 - 999
		case n < 6:
			data[3] += 1 // 1,000 - 999,999
		default:
			data[4] += 1 // 1,000,000+
		}
	}

	return data
}

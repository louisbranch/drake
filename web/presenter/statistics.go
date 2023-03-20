package presenter

import (
	"math"

	"github.com/louisbranch/drake"
	"golang.org/x/text/message"
)

type Statistics struct {
	Survey
	Labels [][]string
	values [][]float64
	Suveys []drake.Survey
}

func NewStatistics(surveys []drake.Survey, printer *message.Printer) Statistics {
	s := Statistics{
		Survey: Survey{Printer: printer},
		Suveys: surveys,
	}

	s.Labels, s.values = s.extract()

	return s
}

func (s Statistics) Data() [][]int {
	data := make([][]int, len(s.values))

	for i := range s.values {
		data[i] = make([]int, len(s.values[i]))
	}

	for _, survey := range s.Suveys {
		if i, ok := s.match(s.values[0], survey.PresurveyAssessment); ok {
			data[0][i] += 1
		}

		if i, ok := s.match(s.values[1], survey.R); ok {
			data[1][i] += 1
		}

		if i, ok := s.match(s.values[2], survey.Fp); ok {
			data[2][i] += 1
		}

		if i, ok := s.match(s.values[3], survey.Ne); ok {
			data[3][i] += 1
		}

		if i, ok := s.match(s.values[4], survey.Fl); ok {
			data[4][i] += 1
		}

		if i, ok := s.match(s.values[5], survey.Fi); ok {
			data[5][i] += 1
		}

		if i, ok := s.match(s.values[6], survey.Fc); ok {
			data[6][i] += 1
		}

		if i, ok := s.match(s.values[7], survey.L); ok {
			data[7][i] += 1
		}

		if survey.PostsurveyLearnGain != nil {
			if *survey.PostsurveyLearnGain {
				data[8][0] += 1
			} else {
				data[8][1] += 1
			}
		}

	}

	return data
}

func (s Statistics) match(vals []float64, v *float64) (int, bool) {
	if v == nil {
		return -1, false
	}

	for i, f := range vals {
		if math.Abs(f-*v) < 0.1 {
			return i, true
		}
	}

	return -1, false
}

func (s Statistics) extract() ([][]string, [][]float64) {
	labels := [][]string{}
	values := [][]float64{}

	l, v := s.collect(s.Survey.NValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.RValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.FpValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.NeValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.FlValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.FiValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.FcValues())
	labels = append(labels, l)
	values = append(values, v)

	l, v = s.collect(s.Survey.LValues())
	labels = append(labels, l)
	values = append(values, v)

	labels = append(labels, []string{
		s.Printer.Sprintf("Yes"),
		s.Printer.Sprintf("No"),
	})

	values = append(values, []float64{1, 0})

	return labels, values
}

func (s Statistics) collect(opts []SurveyOption) ([]string, []float64) {
	labels := make([]string, len(opts))
	values := make([]float64, len(opts))
	for i, opt := range opts {
		text := opt.Text
		if len(text) > 20 {
			labels[i] = text[:20] + "..."
		} else {
			labels[i] = text
		}

		values[i] = opt.Value
	}
	return labels, values
}

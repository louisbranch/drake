package presenter

import (
	"math"

	"github.com/louisbranch/drake"
	"golang.org/x/text/message"
)

// Source: https://doi.org/10.48550/arXiv.1301.6411
// Milky Way galactic disk cylinder in kpc
const MILKY_WAY_R_KPC = 12
const MILKY_WAY_H_KPC = 1

const KPC_TO_LY = 3262

// Number of stars in the Milky Way: 100 billion
const MILKY_WAY_STARS = 1e11

// Age of the Milky Way: 13.61 billion years
const MILKY_WAY_AGE = 13.61e9

type Survey struct {
	drake.Survey
	*message.Printer
}

type SurveyOption struct {
	Text  string
	Value float64
}

func (s Survey) Estimation() string {
	estimation := 0
	if s.N != nil {
		estimation = int(*s.N)
	}
	return s.Printer.Sprintf("You have estimated that there are %d civilizations in the Milky Way.", estimation)
}

func (s Survey) Prediction() string {
	prediction := 0
	if s.PresurveyAssessment != nil {
		prediction = int(*s.PresurveyAssessment)
	}
	return s.Printer.Sprintf("You have predicted that there are %d civilizations in the Milky Way.", prediction)
}

func (s Survey) Difference() string {
	prediction := 0
	if s.PresurveyAssessment != nil {
		prediction = int(*s.PresurveyAssessment)
	}
	return s.Printer.Sprintf("This is a difference of %d orders of magnitude from your initial prediction of %d.", s.difference(), prediction)
}

func (s Survey) difference() int64 {
	if s.N == nil || s.PresurveyAssessment == nil {
		return 0
	}

	n := math.Log10(*s.N)
	if n < 1 {
		n = 0
	}
	a := math.Log10(*s.PresurveyAssessment)
	if a < 1 {
		a = 0
	}
	diff := math.Floor(math.Abs(n - a))
	return int64(diff)
}

func (s Survey) NQuestion() string {
	return s.Printer.Sprintf("How many technological advanced civilizations exist in the Milky Way?")
}

func (s Survey) NValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Only us, we are all alone"),
			Value: 0,
		},
		{
			Text:  s.Printer.Sprintf("A few dozens"),
			Value: 10,
		},
		{
			Text:  s.Printer.Sprintf("A few hundreds"),
			Value: 100,
		},
		{
			Text:  s.Printer.Sprintf("A few thousands"),
			Value: 1e3,
		},
		{
			Text:  s.Printer.Sprintf("A few millions"),
			Value: 1e6,
		},
	}
}

func (s Survey) RQuestion() string {
	return s.Printer.Sprintf("About how many stars are there in our galaxy?")
}

func (s Survey) RValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Around 100 million"),
			Value: 1e8 / MILKY_WAY_AGE,
		},
		{
			Text:  s.Printer.Sprintf("Around 500 million"),
			Value: 5e8 / MILKY_WAY_AGE,
		},
		{
			Text:  s.Printer.Sprintf("Around 100 billion"),
			Value: 1e11 / MILKY_WAY_AGE,
		},
		{
			Text:  s.Printer.Sprintf("Around 500 billion"),
			Value: 5e11 / MILKY_WAY_AGE,
		},
		{
			Text:  s.Printer.Sprintf("More than 1 trillion"),
			Value: 1e12 / MILKY_WAY_AGE,
		},
	}
}

func (s Survey) FpQuestion() string {
	return s.Printer.Sprintf("Approximately what percentage of the stars in our galaxy have planets in orbit?")
}

func (s Survey) FpValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Very few stars have planets in orbit"),
			Value: 0.01,
		},
		{
			Text:  s.Printer.Sprintf("20%%"),
			Value: 0.2,
		},
		{
			Text:  s.Printer.Sprintf("40%%"),
			Value: 0.4,
		},
		{
			Text:  s.Printer.Sprintf("60%%"),
			Value: 0.6,
		},
		{
			Text:  s.Printer.Sprintf("80%%-90%%"),
			Value: 0.85,
		},
	}
}

func (s Survey) NeQuestion() string {
	return s.Printer.Sprintf("On average, how many habitable (Earth-like) planets are there per planetary system?")
}

func (s Survey) NeValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Very few (not even one in every solar system)"),
			Value: 0.1,
		},
		{
			Text:  s.Printer.Sprintf("1"),
			Value: 1,
		},
		{
			Text:  s.Printer.Sprintf("2 or 3"),
			Value: 2.5,
		},
		{
			Text:  s.Printer.Sprintf("4 or 5"),
			Value: 4.5,
		},
		{
			Text:  s.Printer.Sprintf("More than 5"),
			Value: 5.5,
		},
	}
}

func (s Survey) FlQuestion() string {
	return s.Printer.Sprintf("On what fraction of habitable planets will life develop?")
}

func (s Survey) FlValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("A very small percentage, life is very rare"),
			Value: 0.001,
		},
		{
			Text:  s.Printer.Sprintf("20%%"),
			Value: 0.2,
		},
		{
			Text:  s.Printer.Sprintf("50%%"),
			Value: 0.5,
		},
		{
			Text:  s.Printer.Sprintf("80%%"),
			Value: 0.8,
		},
		{
			Text:  s.Printer.Sprintf("100%%%%, if conditions are favorable, life is inevitable"),
			Value: 1,
		},
	}
}

func (s Survey) FiQuestion() string {
	return s.Printer.Sprintf("On what fraction of planets on which life develops will life evolve to intelligence?")
}

func (s Survey) FiValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("A very small percentage"),
			Value: 0.001,
		},
		{
			Text:  s.Printer.Sprintf("1%%-10%%"),
			Value: 0.05,
		},
		{
			Text:  s.Printer.Sprintf("50%%"),
			Value: 0.5,
		},
		{
			Text:  s.Printer.Sprintf("80%%"),
			Value: 0.8,
		},
		{
			Text:  s.Printer.Sprintf("100%%"),
			Value: 1,
		},
	}
}

func (s Survey) FcQuestion() string {
	return s.Printer.Sprintf("What fraction of planets on which life evolves to intelligence will the intelligence develop a technological civilization capable of radio communication?")
}

func (s Survey) FcValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Less than 1%%%%"),
			Value: 0.001,
		},
		{
			Text:  s.Printer.Sprintf("About 10%%%%"),
			Value: 0.1,
		},
		{
			Text:  s.Printer.Sprintf("20%%-30%%"),
			Value: 0.25,
		},
		{
			Text:  s.Printer.Sprintf("50%%-60%%"),
			Value: 0.55,
		},
		{
			Text:  s.Printer.Sprintf("100%%%%, a technological civilization will always eventually develop if life is present"),
			Value: 1,
		},
	}
}

func (s Survey) LQuestion() string {
	return s.Printer.Sprintf("What is the average lifetime of a technological civilization capable of communication?")
}

func (s Survey) LValues() []SurveyOption {
	values := []SurveyOption{}
	for _, n := range []int{100, 1000, 10000, 100000} {
		values = append(values, SurveyOption{
			Text:  s.Printer.Sprintf("%d years", n),
			Value: float64(n),
		})

	}
	values = append(values, SurveyOption{
		Text:  s.Printer.Sprintf("%d years or more", 1000000),
		Value: 1000000,
	})
	return values

}

func (s Survey) LearnGainQuestion() string {
	return s.Printer.Sprintf("Do you agree with your most recent estimation?")
}

func (s Survey) LearnGainValues() []SurveyOption {
	return []SurveyOption{
		{
			Text:  s.Printer.Sprintf("Yes, I am more confident using the Drake Equation calculation"),
			Value: 1,
		},
		{
			Text:  s.Printer.Sprintf("Yes, the Drake Equation includes factors I didn't think during my initial prediction"),
			Value: 1,
		},
		{
			Text:  s.Printer.Sprintf("No, I am more confident with my initial prediction"),
			Value: 0,
		},
		{
			Text:  s.Printer.Sprintf("No, any guess is good as mine"),
			Value: 0,
		},
	}
}

func (s Survey) AvgPredictionDistance() string {
	d, ok := s.avgDistance(s.PresurveyAssessment)
	if !ok {
		return s.Printer.Sprintf("There would be no other civilization in the galaxy.")
	}
	return s.Printer.Sprintf("The average distance to another civilization would be %d light-years.", d)
}

func (s Survey) AvgEstimationDistance() string {
	d, ok := s.avgDistance(s.N)
	if !ok {
		return s.Printer.Sprintf("There would be no other civilization in the galaxy.")
	}
	return s.Printer.Sprintf("The average distance to another civilization would be %d light-years.", d)
}

func (s Survey) avgDistance(val *float64) (int, bool) {
	if val == nil {
		return 0, false
	}

	n := *val

	if n < 1 {
		return 0, false
	}

	// Source: https://doi.org/10.48550/arXiv.1301.6411
	if n < 1000 {
		return int(2 * MILKY_WAY_R_KPC / math.Sqrt(n) * KPC_TO_LY), true
	}

	v := math.Pow(MILKY_WAY_R_KPC, 2) * MILKY_WAY_H_KPC * math.Pi

	return int(2 * math.Pow(3*v/(4*math.Pi*n), 1.0/3) * KPC_TO_LY), true
}

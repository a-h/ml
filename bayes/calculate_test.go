package bayes

import "testing"

func TestProbability(t *testing.T) {
	categoryDog, categoryCat := Category(1), Category(2)

	categoryBrownHair, categoryBlondeHair := Category(1), Category(2)
	categoryBlueEyes, categoryBrownEyes := Category(3), Category(4)

	tests := []struct {
		name     string
		data     Data
		is       Category
		given    []Category
		expected float64
	}{
		{
			name: "dog or cat",
			data: []Datum{
				NewDatum("Felix", categoryCat),
				NewDatum("Shep", categoryDog),
			},
			is:       categoryCat,
			given:    []Category{},
			expected: 0.5,
		},
		{
			name: "brown or blue eyes",
			data: []Datum{
				NewDatum("John", categoryBlueEyes, categoryBlondeHair),
				NewDatum("Jane", categoryBlueEyes, categoryBlondeHair),
				NewDatum("Janet", categoryBlueEyes, categoryBrownHair),
				NewDatum("June", categoryBrownEyes, categoryBrownHair),
			},
			is:       categoryBlondeHair,
			given:    []Category{categoryBlueEyes},
			expected: float64(2) / float64(3),
		},
		{
			name: "dog, cat or lizard",
			data: []Datum{
				NewDatum("Felix", categoryCat),
				NewDatum("Shep", categoryDog),
			},
			is:       Category(3),             // Lizard
			given:    []Category{Category(4)}, // Orangutang
			expected: 0.0,
		},
	}

	for _, test := range tests {
		actual := test.data.Probability(test.is, test.given...)
		if test.expected != actual {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

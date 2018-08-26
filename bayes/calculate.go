package bayes

// Category of the data.
type Category int

// NewDatum creates new data.
func NewDatum(v interface{}, categories ...Category) Datum {
	d := Datum{
		Value:      v,
		Categories: make(map[Category]interface{}, len(categories)),
	}
	for _, c := range categories {
		d.Categories[c] = struct{}{}
	}
	return d
}

// Datum and its associated categories.
type Datum struct {
	Value      interface{}
	Categories map[Category]interface{}
}

// Data is a collection.
type Data []Datum

// Probability of the data being in a category.
func (d Data) Probability(is Category, given ...Category) float64 {
	filtered := filter(d, given...)
	if len(filtered) == 0 {
		return 0
	}
	matchesInFilteredData := filter(filtered, is)
	return float64(len(matchesInFilteredData)) / float64(len(filtered))
}

func filter(d Data, categories ...Category) (filtered Data) {
	for _, dm := range d {
		if !matches(dm, categories...) {
			continue
		}
		filtered = append(filtered, dm)
	}
	return
}

func matches(dm Datum, categories ...Category) bool {
	for _, c := range categories {
		if _, ok := dm.Categories[c]; !ok {
			return false
		}
	}
	return true
}

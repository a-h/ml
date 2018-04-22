package training

import (
	"testing"
)

func TestCompleteUsesAllTrainingData(tt *testing.T) {
	t := &traineeMock{}
	d := []Data{
		{
			Input:    []float64{0},
			Expected: []float64{0},
		},
		{
			Input:    []float64{1},
			Expected: []float64{1},
		},
	}
	a := &algorithmMock{}

	distanceCalled := 0
	dist := func(p []float64, q []float64) (d float64, err error) {
		distanceCalled++
		return 0.0, nil
	}

	iterationCount := 2

	Complete(t, d, a, dist, StopAfterXIterations(iterationCount))

	expectedCalculations := len(d) * iterationCount
	if t.calculateCalled != expectedCalculations {
		tt.Errorf("expected %d training calculations, but %d were carried out", expectedCalculations, t.calculateCalled)
	}
	expectedDistanceCalculations := expectedCalculations
	if distanceCalled != expectedDistanceCalculations {
		tt.Errorf("expected %d distance calculations, but %d were carried out", expectedDistanceCalculations, distanceCalled)
	}
	expectedNextCalled := len(d)
	if a.nextCalled != expectedNextCalled {
		tt.Errorf("expected %d next called, but %d were carried out", expectedNextCalled, a.nextCalled)
	}
	expectedSetMemoryCalled := len(d)
	if t.setMemoryCalled != expectedSetMemoryCalled {
		tt.Errorf("expected %d set memory called, but %d were carried out", expectedSetMemoryCalled, t.setMemoryCalled)
	}
	if a.bestErrorCalled != 2 {
		tt.Errorf("expected bestMemory to be called twice, but was called %d", a.bestErrorCalled)
	}
}

type traineeMock struct {
	calculateCalled     int
	getMemorySizeCalled int
	getMemoryCalled     int
	setMemoryCalled     int
	memory              []float64
}

func (tm *traineeMock) Calculate(input []float64) (output []float64, err error) {
	tm.calculateCalled++
	return input, nil
}
func (tm *traineeMock) GetMemorySize() int {
	tm.getMemorySizeCalled++
	return len(tm.memory)
}
func (tm *traineeMock) GetMemory() []float64 {
	tm.getMemoryCalled++
	return tm.memory
}
func (tm *traineeMock) SetMemory(m []float64) {
	tm.setMemoryCalled++
	tm.memory = m
}

type algorithmMock struct {
	nextCalled       int
	bestErrorCalled  int
	bestMemoryCalled int
	memory           []float64
	e                float64
}

func (am *algorithmMock) Next(e float64) []float64 {
	am.nextCalled++
	am.e = e
	return am.memory
}
func (am *algorithmMock) BestMemory() (memory []float64) {
	am.bestMemoryCalled++
	return am.memory
}
func (am *algorithmMock) BestError() (e float64) {
	am.bestErrorCalled++
	return am.e
}

func TestStopAfterXIterations(t *testing.T) {
	tests := []struct {
		name       string
		stopper    Stopper
		iterations int
		e          float64
		expected   bool
	}{
		{
			name:       "Don't stop until iterations have been cleared",
			stopper:    StopAfterXIterations(10),
			iterations: 9,
			e:          1.0,
			expected:   false,
		},
		{
			name:       "Stop when iterations have been cleared",
			stopper:    StopAfterXIterations(10),
			iterations: 10,
			e:          1.0,
			expected:   true,
		},
	}

	for _, test := range tests {
		actual := test.stopper(test.iterations, test.e)
		if actual != test.expected {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestStopWhenErrorIsLessThan(t *testing.T) {
	tests := []struct {
		name       string
		stopper    Stopper
		iterations int
		e          float64
		expected   bool
	}{
		{
			name:       "Don't stop until error is less than the value",
			stopper:    StopWhenErrorIsLessThan(1.0),
			iterations: 9,
			e:          2.0,
			expected:   false,
		},
		{
			name:       "Stop when error is less than the value",
			stopper:    StopWhenErrorIsLessThan(1.0),
			iterations: 10,
			e:          0.9,
			expected:   true,
		},
	}

	for _, test := range tests {
		actual := test.stopper(test.iterations, test.e)
		if actual != test.expected {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestStopWhenErrorIsGreaterThan(t *testing.T) {
	tests := []struct {
		name       string
		stopper    Stopper
		iterations int
		e          float64
		expected   bool
	}{
		{
			name:       "Don't stop until error is greater than the value",
			stopper:    StopWhenErrorIsGreaterThan(1.0),
			iterations: 9,
			e:          0.5,
			expected:   false,
		},
		{
			name:       "Stop when error is greater than the value",
			stopper:    StopWhenErrorIsGreaterThan(1.0),
			iterations: 10,
			e:          1.5,
			expected:   true,
		},
	}

	for _, test := range tests {
		actual := test.stopper(test.iterations, test.e)
		if actual != test.expected {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestStopWhenChannelReceives(t *testing.T) {
	s, c := StopWhenChannelReceives()
	beforeSent := s(0, 1.0)
	if beforeSent {
		t.Fatalf("expected to carry on until the channel receives an item")
	}
	c <- struct{}{}
	afterSent := s(1, 2.0)
	if !afterSent {
		t.Fatalf("expected to stop after the channel receives an item")
	}
}

package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg is zero  - expected error not be nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second arg is zero  - expected error not be nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestEstimateValueSmall(t *testing.T) {
	result := EstimateValue(9)
	assert.Equal(t, "small", result)
}

func TestEstimateValueMedium(t *testing.T) {
	result := EstimateValue(99)
	assert.Equal(t, "medium", result)
}

func TestEstimateValueBig(t *testing.T) {
	result := EstimateValue(100)
	assert.Equal(t, "big", result)
}

func TestEstimateValueTableDriven(t *testing.T) {
	testCases := []struct {
		Name          string
		InputValue    int
		ExpectedValue string
	}{
		{
			Name:          "Small",
			InputValue:    9,
			ExpectedValue: "small",
		},
		{
			Name:          "Medium",
			InputValue:    99,
			ExpectedValue: "medium",
		},
		{
			Name:          "Big",
			InputValue:    100,
			ExpectedValue: "big",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedValue, EstimateValue(tc.InputValue))
		})
	}
}

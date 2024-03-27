package data

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/uuid"
)

func TestValidator_Required(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"required", nil, true},
		{"required", float64(1), false},
		{"required", 1.2, false},
		{"required", 3.2345679, false},
		{"required", true, false},
		{"required", false, false},
		{"required", "", false},
		{"required", float64(0), false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Required(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

func TestValidator_Type(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"string", "", false},
		{"string", "test", false},
		{"string", "2", false},
		{"string", "2.2", false},
		{"string", nil, true},
		{"string", 2, true},
		{"string", math.MaxFloat64, true},
		{"string", true, true},
		{"number", float64(2), false},
		{"number", float64(math.MaxInt64), false},
		{"number", math.MaxFloat64, false},
		{"number", 2.27892, false},
		{"number", 2, true},
		{"number", true, true},
		{"number", "2.2345", true},
		{"boolean", true, false},
		{"boolean", false, false},
		{"boolean", float64(1), true},
		{"boolean", float64(0), true},
		{"boolean", 0, true},
		{"boolean", 1, true},
		{"boolean", "true", true},
		{"boolean", "false", true},
		{"boolean", nil, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Type(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

func TestValidator_Min(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"min=1", float64(2), false}, // Unmarshalling JSON produces floats
		{"min=3", float64(1), true},
		{"min=1.1", 1.2, false},
		{"min=3", 1.2, true},
		{"min=3.23456789", 3.23456789, true},
		{"min=3.23456789", 3.2345679, false},
		{"min=5", "words", true},
		{"min=4", "words", false},
		{"min=2", []string{"words"}, true},
		{"min=2", []string{"words", "words", "words"}, false},
		{"min=3", nil, true},
		{"min=3", 4, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Min(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

func TestValidator_Max(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"max=1", float64(2), true}, // Unmarshalling JSON produces floats
		{"max=3", float64(1), false},
		{"max=1.1", 1.2, true},
		{"max=3", 1.2, false},
		{"max=3.23456789", 3.23456789, true},
		{"max=3.23456789", 3.2345679, true},
		{"max=3.23456789", 3.234567889, false},
		{"max=5", "words", true},
		{"max=5", "word", false},
		{"max=2", []string{"words"}, false},
		{"max=2", []string{"words", "words", "words"}, true},
		{"max=3", nil, true},
		{"max=3", 2, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Max(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

func TestValidator_Int(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"int", float64(2), false}, // Unmarshalling JSON produces floats
		{"int", 1.2, true},
		{"int", 3.23456789, true},
		{"int", float64(math.MinInt64), true}, // Should fail - precision
		{"int", float64(math.MaxInt64), true}, // Should fail - precision
		{"int", math.MaxFloat64, true},
		{"int", "words", true},
		{"int", nil, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Int(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

func TestValidator_UUID(t *testing.T) {
	var tests = []struct {
		rule    string
		value   any
		wantErr bool
	}{
		{"uuid", "6d61a221-4a50-45a0-8421-cd3609bc5526", false},
		{"uuid", "6d61a221-4a50-35a0-8421-cd3609bc5526", true},
		{"uuid", uuid.Nil, true},
		{"uuid", float64(2), true},
		{"uuid", 1.2, true},
		{"uuid", "words", true},
		{"uuid", nil, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.rule, tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := UUID(tt.rule, tt.value)
			if ans != nil && !tt.wantErr {
				t.Errorf("got %v, wantErr %v", ans, tt.wantErr)
			}
		})
	}
}

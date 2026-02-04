package src

import "testing"

func TestAbs_PositiveNumber(t *testing.T) {
	result := Abs(5)
	expected := 5
	if result != expected {
		t.Errorf("Abs(5) = %d; want %d", result, expected)
	}
}

func TestAbs_NegativeNumber(t *testing.T) {
	result := Abs(-5)
	expected := 5
	if result != expected {
		t.Errorf("Abs(-5) = %d; want %d", result, expected)
	}
}

func TestAbs_Zero(t *testing.T) {
	result := Abs(0)
	expected := 0
	if result != expected {
		t.Errorf("Abs(0) = %d; want %d", result, expected)
	}
}

func TestAbs_LargePositiveNumber(t *testing.T) {
	result := Abs(1000000)
	expected := 1000000
	if result != expected {
		t.Errorf("Abs(1000000) = %d; want %d", result, expected)
	}
}

func TestAbs_LargeNegativeNumber(t *testing.T) {
	result := Abs(-1000000)
	expected := 1000000
	if result != expected {
		t.Errorf("Abs(-1000000) = %d; want %d", result, expected)
	}
}

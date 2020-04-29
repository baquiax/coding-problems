package solution

import (
	"fmt"
	"testing"
)

func TestValidNumber(t *testing.T) {
	value := "IX"

	_, error := convertRomanToDecimalNotation(value)

	if error != nil {
		t.Error(error)
	}
}

func TestInvalidNumber(t *testing.T) {
	value := "cosa"

	_, error := convertRomanToDecimalNotation(value)

	if error == nil {
		t.Error(fmt.Errorf("%s is not a valid number", value))
	}
}

func BenchmarkLengthTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		value := "IX"
		convertRomanToDecimalNotation(value)
	}
}

func BenchmarkLengthFive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		value := "MDCVI"
		convertRomanToDecimalNotation(value)
	}
}

func BenchmarkNotImmediateError(b *testing.B) {
	for n := 0; n < b.N; n++ {
		value := "hello"
		convertRomanToDecimalNotation(value)
	}
}

func BenchmarkImmediateError(b *testing.B) {
	for n := 0; n < b.N; n++ {
		value := "MDCXXXVIZ"
		convertRomanToDecimalNotation(value)
	}
}

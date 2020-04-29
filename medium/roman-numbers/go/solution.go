package solution

import (
	"fmt"
)

// RomanNumberValues Representation of the numbers in Roman's notation
var RomanNumberValues = map[byte]int{
	'M': 1000,
	'm': 1000,
	'D': 500,
	'd': 500,
	'C': 100,
	'c': 100,
	'L': 50,
	'l': 50,
	'X': 10,
	'x': 10,
	'V': 5,
	'v': 5,
	'I': 1,
	'i': 1,
}

func convertRomanToDecimalNotation(input string) (int, error) {
	var result = 0
	var lastValue = 0

	for i := 0; i < len(input); i++ {
		currentLetter := input[i]

		currentValue, error := getLetterValue(currentLetter)
		if error != nil {
			return 0, fmt.Errorf("Error at position %d: %s", i+1, error)
		}

		if lastValue < currentValue {
			result += (currentValue - 2*lastValue)
		} else {
			result += currentValue
		}

		lastValue = currentValue
	}

	return result, nil
}

func getLetterValue(letter byte) (int, error) {
	value := RomanNumberValues[letter]
	if value == 0 {
		return 0, fmt.Errorf("%c is not a valid Roman number", letter)
	}

	return value, nil
}

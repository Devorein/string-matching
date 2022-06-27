package main

import (
	"fmt"
)

func generatePiTable(pattern string) []int{
	piTable := []int{0}
	patternLength := len(pattern)

	if patternLength == 1 {
		return piTable
	}

	i := 0
	j := 1

	for j < patternLength {

		if pattern[i] == pattern[j] {
			piTable = append(piTable, i + 1)
			i += 1
		} else {
			i = 0
			piTable = append(piTable, 0)
		}

		j += 1
	}

	return piTable
}

func kmp (inputString string, pattern string) bool {
	inputStringLength:= len(inputString)
	patternLength := len(pattern)

	if patternLength > inputStringLength {
		return false
	}

	inputStringPointer := 0
	patternPointer := -1

	piTable := generatePiTable(pattern)

	for inputStringPointer < inputStringLength {
		patternChar := pattern[patternPointer + 1]
		inputStringChar := inputString[inputStringPointer]
		if patternChar == inputStringChar {
			patternPointer += 1
			inputStringPointer += 1

			if patternPointer + 1 == patternLength {
				return true
			}
		} else if patternPointer == -1 {
			inputStringPointer += 1
		} else {
			patternPointer = piTable[patternPointer] - 1
			if patternPointer + 1 == patternLength {
				return true
			}
		}
	}

	return patternPointer + 1 == patternLength
}

func main() {
	pattern:= "ababd"
	inputString := "abababd"

	fmt.Println(kmp(inputString, pattern))
}
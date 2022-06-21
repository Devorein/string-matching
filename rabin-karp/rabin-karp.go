package main

import (
	"fmt"
	"math"
)

var charAsciiDict = map[string]int{
	"0": 48,
	"1": 49,
	"2": 50,
	"3": 51,
	"4": 52,
	"5": 53,
	"6": 54,
	"7": 55,
	"8": 56,
	"9": 57,
	"!": 33,
	"\"": 34,
	"#": 35,
	"$": 36,
	"%": 37,
	"&": 38,
	"'": 39,
	"(": 40,
	")": 41,
	"*": 42,
	"+": 43,
	",": 44,
	"-": 45,
	".": 46,
	"/": 47,
	":": 58,
	";": 59,
	"<": 60,
	"=": 61,
	">": 62,
	"?": 63,
	"@": 64,
	"A": 65,
	"B": 66,
	"C": 67,
	"D": 68,
	"E": 69,
	"F": 70,
	"G": 71,
	"H": 72,
	"I": 73,
	"J": 74,
	"K": 75,
	"L": 76,
	"M": 77,
	"N": 78,
	"O": 79,
	"P": 80,
	"Q": 81,
	"R": 82,
	"S": 83,
	"T": 84,
	"U": 85,
	"V": 86,
	"W": 87,
	"X": 88,
	"Y": 89,
	"Z": 90,
	"[": 91,
	"\\": 92,
	"]": 93,
	"^": 94,
	"_": 95,
	"`": 96,
	"a": 97,
	"b": 98,
	"c": 99,
	"d": 100,
	"e": 101,
	"f": 102,
	"g": 103,
	"h": 104,
	"i": 105,
	"j": 106,
	"k": 107,
	"l": 108,
	"m": 109,
	"n": 110,
	"o": 111,
	"p": 112,
	"q": 113,
	"r": 114,
	"s": 115,
	"t": 116,
	"u": 117,
	"v": 118,
	"w": 119,
	"x": 120,
	"y": 121,
	"z": 122,
	"{": 123,
	"|": 124,
	"}": 125,
	"~": 126,
}

func calculateRollingHash(oldHash int, firstCharOfOldSubstring string, lastCharOfOldSubstring string, patternLength int, base int) int {
	return int(math.Floor((float64(oldHash) - float64(charAsciiDict[firstCharOfOldSubstring])) / float64(base)))+ charAsciiDict[lastCharOfOldSubstring] * int(math.Pow(float64(base), float64(patternLength - 1)))
}

func calculateHash(substring string, base int) int {
	substringLength:= len(substring)
	hash:= 0
	for idx:=0; idx < substringLength; idx+=1 {
		hash += int(float64(charAsciiDict[string(substring[idx])]) * math.Pow(float64(base), float64(idx)))
	}

	return hash
}

func matchStrings(substring string, pattern string) bool {
	substringLength:= len(substring)
	patternLength := len(pattern)

	if (substringLength != patternLength) {
		return false
	}

	for idx:= 0; idx < substringLength; idx += 1 {
		if substring[idx] != pattern[idx] {
			return false
		}
	}

	return true
}

func rabinKarp(inputString string, pattern string) bool {
	inputStringLength:= len(inputString)
	patternLength:= len(pattern)

	base := 3

	if (patternLength > inputStringLength) {
		return false
	}

	patternHash:= calculateHash(pattern, base)
	currentIdx:= patternLength
	substring:= inputString[0:currentIdx]
	currentHash:= calculateHash(substring, base)

	for currentIdx <= inputStringLength {
		if patternHash == currentHash {
			stringMatches:= matchStrings(pattern, substring)
			if stringMatches {
				return true
			}
		}

		if (currentIdx == inputStringLength) {
			return false;
		}

		firstCharOfOldSubstring:= substring[0]
		substring = inputString[(currentIdx + 1) - patternLength : currentIdx + 1]
		currentIdx = currentIdx + 1
		currentHash = calculateRollingHash(currentHash, string(firstCharOfOldSubstring), string(substring[len(substring) - 1]), patternLength, base)
	}

	return false
}

func main() {
	pattern:= "abc"
	inputString:= "abaabac"
	fmt.Println(rabinKarp(inputString, pattern))
}
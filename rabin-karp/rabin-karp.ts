const CharAsciiRecord: Record<string, number> = {
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
  "~": 126
}

function calculateRollingHash(oldHash: number, firstCharOfOldSubstring: string, lastCharOfNewSubstring: string, patternLength: number, base: number) {
  return Math.floor((oldHash - CharAsciiRecord[firstCharOfOldSubstring]) / base) + CharAsciiRecord[lastCharOfNewSubstring] * Math.pow(base, patternLength - 1)
}

function calculateHash(substring: string, base: number) {
  const substringLength = substring.length;
  let hash = 0;
  for (let idx = 0; idx < substringLength; idx++) {
    hash += CharAsciiRecord[substring[idx]] * Math.pow(base, idx);
  }

  return hash;
}

function matchStrings(substring: string, pattern: string) {
  const substringLength = substring.length;
  const patternLength = pattern.length;
  if (substringLength !== patternLength) {
    return false;
  }

  for (let idx = 0; idx < substringLength; idx += 1) {
    if (substring[idx] !== pattern[idx]) {
      return false;
    }
  }
  return true;
}

function rabinKarp(inputString: string, pattern: string) {
  const patternLength = pattern.length;
  const inputStringLength = inputString.length;

  const base = 3;

  if (patternLength > inputStringLength) {
    return false;
  }

  const patternHash = calculateHash(pattern, base);
  let currentIdx = patternLength;
  let substring = inputString.slice(0, currentIdx);
  let currentHash = calculateHash(substring, base);

  while (currentIdx < inputStringLength) {
    if (patternHash === currentHash) {
      const stringsMatched = matchStrings(pattern, substring)
      if (stringsMatched) {
        return true
      }
    }
    const oldSubstring = substring;
    substring = inputString.slice((currentIdx + 1) - patternLength, currentIdx += 1);
    currentHash = calculateRollingHash(currentHash, oldSubstring[0], substring[substring.length - 1], patternLength, base)
  }

  if (patternHash === currentHash) {
    const stringsMatched = matchStrings(pattern, substring)
    if (stringsMatched) {
      return true
    }
  }

  return false;
}

const pattern = "abc";
const inputString = "defghiabacabac";

console.log(rabinKarp(inputString, pattern))
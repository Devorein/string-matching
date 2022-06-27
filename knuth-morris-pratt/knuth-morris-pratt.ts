function generatePiTable(pattern: string) {
  const piTable: number[] = [0]

  if (pattern.length === 1) {
    return piTable
  }

  let i = 0;
  let j = 1;

  const stringLength = pattern.length;

  while (j < stringLength) {
    if (pattern[i] === pattern[j]) {
      piTable.push(i + 1)
      i += 1;
    } else {
      i = 0;
      piTable.push(0)
    }
    j += 1;
  }

  return piTable;
}

function kmp(inputString: string, pattern: string) {
  const inputStringLength = inputString.length;
  const patternLength = pattern.length;
  if (patternLength > inputStringLength) {
    return false;
  }

  let inputStringPointer = 0;
  let patternPointer = -1;

  const piTable = generatePiTable(pattern);
  while (inputStringPointer < inputStringLength) {
    const patternChar = pattern[patternPointer + 1];
    const inputStringChar = inputString[inputStringPointer]
    if (patternChar === inputStringChar) {
      patternPointer += 1;
      inputStringPointer += 1;
      if (patternPointer + 1 === patternLength) {
        return true;
      }
    } else if (patternPointer === -1) {
      inputStringPointer += 1;
    } else {
      patternPointer = piTable[patternPointer] - 1;
      if (patternPointer + 1 === patternLength) {
        return true;
      }
    }
  }

  return patternPointer + 1 === patternLength;
}

console.log(kmp("ababababbabdcd", "ababd"))
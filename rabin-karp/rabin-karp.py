import math

char_ascii_dict: dict[str, int] = {
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


def calculate_rolling_hash(old_hash: int, first_char_of_old_substring: str, last_char_of_new_substring: str, pattern_length: int, base: int) -> int:
    return math.floor((old_hash - char_ascii_dict[first_char_of_old_substring]) / base) + char_ascii_dict[last_char_of_new_substring] * math.pow(base, pattern_length - 1)


def calculate_hash(substring: str, base: int) -> int:
    substring_length = len(substring)
    hash = 0
    for idx in range(substring_length):
        hash += char_ascii_dict[substring[idx]] * math.pow(base, idx)

    return math.floor(hash)


def match_strings(substring: str, pattern: str) -> bool:
    substring_length = len(substring)
    pattern_length = len(pattern)
    if (substring_length != pattern_length):
        return False

    for idx in range(substring_length):
        if (substring[idx] != pattern[idx]):
            return False
    return True


def rabin_karp(input_string: str, pattern: str) -> bool:
    pattern_length = len(pattern)
    input_string_length = len(input_string)

    base = 3

    if (pattern_length > input_string_length):
        return False

    pattern_hash = calculate_hash(pattern, base)
    current_idx = pattern_length
    substring = input_string[0:current_idx]
    current_hash = calculate_hash(substring, base)

    while (current_idx < input_string_length):
        if (pattern_hash == current_hash):
            strings_matches = match_strings(pattern, substring)
            if (strings_matches):
                return True
        old_substring = substring
        substring = input_string[(current_idx + 1) -
                                 pattern_length: current_idx + 1]
        current_idx = current_idx + 1
        current_hash = calculate_rolling_hash(
            current_hash, old_substring[0], substring[-1], pattern_length, base)

    if (pattern_hash == current_hash):
        strings_matches = match_strings(pattern, substring)
        if (strings_matches):
            return True

    return False


pattern = "abc"
input_string = "defghiabacaba"

print(rabin_karp(input_string, pattern))

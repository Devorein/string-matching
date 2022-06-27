def generate_pi_table(pattern: str):
    pi_table: list[int] = [0]

    if len(pattern) == 1:
        return pi_table

    i = 0
    j = 1

    string_length = len(pattern)

    while j < string_length:
        if pattern[i] == pattern[j]:
            pi_table.append(i + 1)
            i += 1
        else:
            i = 0
            pi_table.append(0)
        j += 1
    return pi_table


def kmp(input_str: str, pattern: str):
    input_str_length = len(input_str)
    pattern_length = len(pattern)
    if (pattern_length > input_str_length):
        return False

    input_str_pointer = 0
    pattern_pointer = -1

    pi_table = generate_pi_table(pattern)
    while input_str_pointer < input_str_length:
        pattern_char = pattern[pattern_pointer + 1]
        input_str_char = input_str[input_str_pointer]
        if pattern_char == input_str_char:
            pattern_pointer += 1
            input_str_pointer += 1
            if pattern_pointer + 1 == pattern_length:
                return True
        elif pattern_pointer == -1:
            input_str_pointer += 1
        else:
            pattern_pointer = pi_table[pattern_pointer] - 1
            if pattern_pointer + 1 == pattern_length:
                return True
    return pattern_pointer + 1 == pattern_length


print(kmp("abababababdcddasdasd", "ababd"))

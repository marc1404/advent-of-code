def parse_instruction(line):
    parts = line.split()

    return [parts[0], int(parts[1])]


def read_instructions(file_path):
    with open(file_path) as file:
        return [parse_instruction(line) for line in file]


def calculate_position(instructions):
    x, z = 0, 0

    for direction, distance in instructions:
        if direction == 'forward':
            x += distance
        elif direction == 'up':
            z -= distance
        elif direction == 'down':
            z += distance
        else:
            assert False, f'Unknown direction: {direction}'

    return x * z


def calculate_position_with_aim(instructions):
    x, z, aim = 0, 0, 0

    for direction, distance in instructions:
        if direction == 'forward':
            x += distance
            z += aim * distance
        elif direction == 'up':
            aim -= distance
        elif direction == 'down':
            aim += distance
        else:
            assert False, f'Unknown direction: {direction}'

    return x * z


example_instructions = read_instructions('example.txt')
input_instructions = read_instructions('input.txt')

print('Part 1')
print(f'Example: {calculate_position(example_instructions)}')
print(f'Input: {calculate_position(input_instructions)}')
print()
print('Part 2')
print(f'Example: {calculate_position_with_aim(example_instructions)}')
print(f'Input: {calculate_position_with_aim(input_instructions)}')

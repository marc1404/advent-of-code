def read_diagnostics_grid(file_path):
    with open(file_path) as file:
        return [parse_diagnostics_line(line.rstrip()) for line in file]


def parse_diagnostics_line(line):
    return [int(bit) for bit in list(line)]


def get_column(grid, column_index):
    column = []

    for row in grid:
        column.append(row[column_index])

    return column


def determine_common_and_uncommon(column):
    zero_count, one_count = 0, 0

    for bit in column:
        if bit == 0:
            zero_count += 1
        elif bit == 1:
            one_count += 1

    return (0, 1) if zero_count > one_count else (1, 0)


def calculate_gamma_and_epsilon(grid):
    gamma_bits, epsilon_bits = [], []

    for column_index, _ in enumerate(grid[0]):
        column = get_column(grid, column_index)
        (gamma_bit, epsilon_bit) = determine_common_and_uncommon(column)

        gamma_bits.append(gamma_bit)
        epsilon_bits.append(epsilon_bit)

    return bits_to_decimal(gamma_bits), bits_to_decimal(epsilon_bits)


def bits_to_decimal(bits):
    bit_string = ''.join(str(bit) for bit in bits)

    return int(bit_string, 2)


def determine_power_consumption(grid):
    (gamma, epsilon) = calculate_gamma_and_epsilon(grid)
    power_consumption = gamma * epsilon

    return power_consumption


def determine_oxygen_and_co2_scrubber_ratings(grid):
    oxygen_grid = grid[:]
    co2_grid = grid[:]

    for column_index, _ in enumerate(grid[0]):
        oxygen_column = get_column(oxygen_grid, column_index)
        co2_column = get_column(co2_grid, column_index)
        (common, _) = determine_common_and_uncommon(oxygen_column)
        (_, uncommon) = determine_common_and_uncommon(co2_column)

        if len(oxygen_grid) > 1:
            oxygen_grid = [row for row in oxygen_grid if row[column_index] == common]

        if len(co2_grid) > 1:
            co2_grid = [row for row in co2_grid if row[column_index] == uncommon]

        if len(oxygen_grid) == 1 and len(co2_grid) == 1:
            break

    oxygen_rating = bits_to_decimal(oxygen_grid[0])
    co2_scrubber_rating = bits_to_decimal(co2_grid[0])

    return oxygen_rating, co2_scrubber_rating


def determine_life_support_rating(grid):
    (oxygen_rating, co2_scrubber_rating) = determine_oxygen_and_co2_scrubber_ratings(grid)
    life_support_rating = oxygen_rating * co2_scrubber_rating

    return life_support_rating


def main():
    example_grid = read_diagnostics_grid('example.txt')
    input_grid = read_diagnostics_grid('input.txt')

    print('Part 1')
    print(f'Example: {determine_power_consumption(example_grid)}')
    print(f'Input: {determine_power_consumption(input_grid)}')
    print('')
    print('Part 2')
    print(f'Example: {determine_life_support_rating(example_grid)}')
    print(f'Input: {determine_life_support_rating(input_grid)}')


if __name__ == '__main__':
    main()

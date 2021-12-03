import sys


def read_measurements(file_path):
    with open(file_path) as file:
        return [int(line.rstrip()) for line in file]


def count_measurement_increases(measurements):
    increase_count = 0
    previous_measurement = sys.maxsize

    for measurement in measurements:
        if measurement > previous_measurement:
            increase_count += 1

        previous_measurement = measurement

    return increase_count


def count_sliding_window_increases(measurements):
    sliding_windows = []

    for i, measurement in enumerate(measurements):
        sliding_window = sum(measurements[i:i + 3])

        sliding_windows.append(sliding_window)

    return count_measurement_increases(sliding_windows)


example_measurements = read_measurements('example.txt')
input_measurements = read_measurements('input.txt')

print('Part 1')
print(f'Example: {count_measurement_increases(example_measurements)}')
print(f'Input: {count_measurement_increases(input_measurements)}')
print()
print('Part 2')
print(f'Example: {count_sliding_window_increases(example_measurements)}')
print(f'Input: {count_sliding_window_increases(input_measurements)}')

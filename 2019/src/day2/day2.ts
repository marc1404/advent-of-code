import consola from 'consola';
import assert from 'assert';
import { day2Input as input } from './input';
import { IntCode } from '../int-code/IntCode';

export function day2(): void {
    const testResult1 = new IntCode([1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50])
        .execute()
        .getIntCode();

    assert.deepStrictEqual(testResult1, [3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50]);

    const testResult2 = new IntCode([1, 0, 0, 0, 99])
        .execute()
        .getIntCode();

    assert.deepStrictEqual(testResult2, [2, 0, 0, 0, 99]);

    const testResult3 = new IntCode([2, 3, 0, 3, 99])
        .execute()
        .getIntCode();

    assert.deepStrictEqual(testResult3, [2, 3, 0, 6, 99]);

    const testResult4 = new IntCode([2, 4, 4, 5, 99, 0])
        .execute()
        .getIntCode();

    assert.deepStrictEqual(testResult4, [2, 4, 4, 5, 99, 9801]);

    const testResult5 = new IntCode([1, 1, 1, 4, 99, 5, 6, 0, 99])
        .execute()
        .getIntCode();

    assert.deepStrictEqual(testResult5, [30, 1, 1, 4, 2, 5, 6, 0, 99]);

    puzzlePart1();
    puzzlePart2();
}

function puzzlePart1(): void {
    const program = [...input];
    program[1] = 12;
    program[2] = 2;

    const intCode = new IntCode(program)
        .execute()
        .getIntCode();

    const [output] = intCode;

    consola.info(`Value at position 0: ${output}`);
}

function puzzlePart2(): void {
    const expected = 19690720;

    for (let noun = 0; noun <= 99; noun++) {
        for (let verb = 0; verb <= 99; verb++) {
            const program = [...input];
            program[1] = noun;
            program[2] = verb;

            const intCode = new IntCode(program)
                .execute()
                .getIntCode();

            const [output] = intCode;

            if (output === expected) {
                consola.info(`The output of ${expected} is produced by noun=${noun} and verb=${verb}`);

                return;
            }
        }
    }
}

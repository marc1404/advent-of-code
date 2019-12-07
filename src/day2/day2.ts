import consola from 'consola';
import assert from 'assert';
import { day2Input as input } from './input';

enum OpCode {
    Addition = 1,
    Multiplication = 2,
    Done = 99
}

export function day2(): void {
    const testResult1 = executeIntCode([1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50]);

    assert.deepStrictEqual(testResult1, [3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50]);

    const testResult2 = executeIntCode([1, 0, 0, 0, 99]);

    assert.deepStrictEqual(testResult2, [2, 0, 0, 0, 99]);

    const testResult3 = executeIntCode([2, 3, 0, 3, 99]);

    assert.deepStrictEqual(testResult3, [2, 3, 0, 6, 99]);

    const testResult4 = executeIntCode([2, 4, 4, 5, 99, 0]);

    assert.deepStrictEqual(testResult4, [2, 4, 4, 5, 99, 9801]);

    const testResult5 = executeIntCode([1, 1, 1, 4, 99, 5, 6, 0, 99]);

    assert.deepStrictEqual(testResult5, [30, 1, 1, 4, 2, 5, 6, 0, 99]);

    puzzlePart1();
    puzzlePart2();
}

function puzzlePart1(): void {
    const program = [...input];
    program[1] = 12;
    program[2] = 2;

    executeIntCode(program);

    const [output] = program;

    consola.info(`Value at position 0: ${output}`);
}

function puzzlePart2(): void {
    const expected = 19690720;

    for (let noun = 0; noun <= 99; noun++) {
        for (let verb = 0; verb <= 99; verb++) {
            const program = [...input];
            program[1] = noun;
            program[2] = verb;

            executeIntCode(program);

            const [output] = program;

            if (output === expected) {
                consola.info(`The output of ${expected} is produced by noun=${noun} and verb=${verb}`);

                return;
            }
        }
    }
}

function executeIntCode(intCode: number[]): number[] {
    let pointer = 0;

    do {
        const [opCode, leftIndex, rightIndex, outputIndex] = intCode.slice(pointer, pointer + 4);

        if (opCode === OpCode.Done) {
            break;
        }

        const operation = getOperation(opCode);
        const left = intCode[leftIndex];
        const right = intCode[rightIndex];
        intCode[outputIndex] = operation(left, right);
        pointer += 4;
    } while (true);

    return intCode;
}

function getOperation(opCode: number): (left: number, right: number) => number {
    if (opCode === OpCode.Addition) {
        return (left: number, right: number) => left + right;
    }

    if (opCode === OpCode.Multiplication) {
        return (left: number, right: number) => left * right;
    }

    throw new Error(`Unknown opcode: ${opCode}!`);
}

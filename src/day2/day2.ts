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

    input[1] = 12;
    input[2] = 2;
    const result = executeIntCode(input);

    consola.info(`Value at position 0: ${result[0]}`);
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

import consola from 'consola';
import assert from 'assert';
import { day5Input } from './input';
import { Instruction } from '../int-code/Instruction';
import { IntCode } from '../int-code/IntCode';

export function day5(): void {
    testInstruction(1002, 2, [0, 1, 0]);
    testInstruction(1101, 1, [1, 1, 0]);
    testInstruction(1, 1, [0, 0, 0]);
    testInstruction(11101, 1, [1, 1, 1]);
    testInstruction(11001, 1, [0, 1, 1]);

    test1();
    test2();
    test3();

    puzzle1();

    test4();
    test5();
    test6();
    test7();
    test8();
    test9();
    test10();

    puzzle2();
}

function testInstruction(instructionValue: number, expectedOpCode: number, expectedParameterNodes: number[]): void {
    const instruction = new Instruction(instructionValue);
    const opCode = instruction.getOpCode();
    const parameterModes = instruction.getParameterModes();

    assert.strictEqual(opCode, expectedOpCode);
    assert.deepStrictEqual(parameterModes, expectedParameterNodes);
}

function test1(): void {
    consola.info('input = output');

    new IntCode([3, 0, 4, 0, 99]).execute();
}

function test2(): void {
    const intCode = new IntCode([1002, 4, 3, 4, 33])
        .execute()
        .getIntCode();

    assert.strictEqual(intCode[4], 99);
}

function test3(): void {
    const intCode = new IntCode([1101, 100, -1, 4, 0])
        .execute()
        .getIntCode();

    assert.strictEqual(intCode[4], 99);
}

function puzzle1(): void {
    consola.info('Please input 1');

    new IntCode(day5Input)
        .execute();
}

function test4(): void {
    consola.info('input = 8 ? 1 : 0');

    new IntCode([3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8])
        .execute();
}

function test5(): void {
    consola.info('input < 8 ? 1 : 0');

    new IntCode([3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8])
        .execute();
}

function test6(): void {
    consola.info('input = 8 ? 1 : 0');

    new IntCode([3, 3, 1108, -1, 8, 3, 4, 3, 99])
        .execute();
}

function test7(): void {
    consola.info('input < 8 ? 1 : 0');

    new IntCode([3, 3, 1107, -1, 8, 3, 4, 3, 99])
        .execute();
}

function test8(): void {
    consola.info('input = 1 ? 1 : 0');

    new IntCode([3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9])
        .execute();
}

function test9(): void {
    consola.info('input = 1 ? 1 : 0');

    new IntCode([3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1])
        .execute();
}

function test10(): void {
    consola.info('input < 8 -> 999 | input = 8 -> 1000 | input > 8 -> 1001');

    new IntCode([3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99])
        .execute();
}

function puzzle2(): void {
    consola.info('Please input 5');

    new IntCode(day5Input)
        .execute();
}

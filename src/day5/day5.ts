import { OpCode } from '../day2/day2';
import consola from 'consola';
import readLineSync from 'readline-sync';
import assert from 'assert';
import { day5Input } from './input';

enum ParameterMode {
    Position = 0,
    Immediate = 1
}

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
    executeIntCode([3, 0, 4, 0, 99]);
}

function test2(): void {
    const intCode = executeIntCode([1002, 4, 3, 4, 33]);

    assert.strictEqual(intCode[4], 99);
}

function test3(): void {
    const intCode = executeIntCode([1101, 100, -1, 4, 0]);

    assert.strictEqual(intCode[4], 99);
}

function puzzle1(): void {
    consola.info('Please input 1');
    executeIntCode(day5Input);
}

function test4(): void {
    consola.info('input = 8 ? 1 : 0');
    executeIntCode([3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8]);
}

function test5(): void {
    consola.info('input < 8 ? 1 : 0');
    executeIntCode([3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8]);
}

function test6(): void {
    consola.info('input = 8 ? 1 : 0');
    executeIntCode([3, 3, 1108, -1, 8, 3, 4, 3, 99]);
}

function test7(): void {
    consola.info('input < 8 ? 1 : 0');
    executeIntCode([3, 3, 1107, -1, 8, 3, 4, 3, 99]);
}

function test8(): void {
    consola.info('input = 1 ? 1 : 0');
    executeIntCode([3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9]);
}

function test9(): void {
    consola.info('input = 1 ? 1 : 0');
    executeIntCode([3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1]);
}

function test10(): void {
    consola.info('input < 8 -> 999 | input = 8 -> 1000 | input > 8 -> 1001');
    executeIntCode([3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99]);
}

function puzzle2(): void {
    consola.info('Please input 5');
    executeIntCode(day5Input);
}

export function executeIntCode(intCode: number[], inputs: number[] = [], outputs: number[] = []): number[] {
    let instructionPointer: number = 0;

    while (true) {
        const instructionValue = intCode[instructionPointer];
        const instruction = new Instruction(instructionValue);

        if (instruction.isDone()) {
            return intCode;
        }

        const opCode = instruction.getOpCode();
        const parameterModes = instruction.getParameterModes();
        instructionPointer = executeOperation(intCode, opCode, instructionPointer, parameterModes, inputs, outputs);
    }
}

function executeOperation(
    intCode: number[],
    opCode: number,
    instructionPointer: number,
    parameterModes: number[],
    inputs: number[],
    outputs: number[]
): number {
    if (opCode === OpCode.Addition) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left + right;

        return instructionPointer + 4;
    }

    if (opCode === OpCode.Multiplication) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left * right;

        return instructionPointer + 4;
    }

    if (opCode === OpCode.Input) {
        const outputPointer = intCode[instructionPointer + 1];
        intCode[outputPointer] = readInput(outputPointer, inputs);

        return instructionPointer + 2;
    }

    if (opCode === OpCode.Output) {
        const outputPointer = intCode[instructionPointer + 1];
        const output = intCode[outputPointer];

        outputs.push(output);
        consola.info(`Value at position ${outputPointer} is ${output}`);

        return instructionPointer + 2;
    }

    if (opCode === OpCode.JumpIfTrue) {
        const a = getParameter(intCode, instructionPointer, parameterModes, 0);
        const b = getParameter(intCode, instructionPointer, parameterModes, 1);

        return a !== 0
            ? b
            : instructionPointer + 3;
    }

    if (opCode === OpCode.JumpIfFalse) {
        const a = getParameter(intCode, instructionPointer, parameterModes, 0);
        const b = getParameter(intCode, instructionPointer, parameterModes, 1);

        return a === 0
            ? b
            : instructionPointer + 3;
    }

    if (opCode === OpCode.LessThan) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left < right ? 1 : 0;

        return instructionPointer + 4;
    }

    if (opCode === OpCode.Equals) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left === right ? 1 : 0;

        return instructionPointer + 4;
    }

    throw new Error(`Unknown opcode: ${opCode}!`);
}

function readInput(outputPointer: number, inputs: number[]): number {
    const input = inputs.shift();

    if (input !== undefined) {
        return input;
    }

    const inputString = readLineSync.question(`Input for position ${outputPointer}: `);

    return Number.parseInt(inputString, 10);
}

function getParameter(intCode: number[], instructionPointer: number, parameterModes: number[], parameterIndex: number): number {
    const parameterMode = parameterModes[parameterIndex];
    const parameterPointer = instructionPointer + 1 + parameterIndex;

    if (parameterMode === ParameterMode.Position) {
        const position = intCode[parameterPointer];

        return intCode[position];
    }

    if (parameterMode === ParameterMode.Immediate) {
        return intCode[parameterPointer];
    }

    throw new Error(`Unknown parameter mode: ${parameterMode}!`);
}

class Instruction {

    private readonly opCode: number;
    private readonly parameterModes: number[];

    constructor(
        private readonly instructionValue: number
    ) {
        const instructionString = instructionValue
            .toString()
            .padStart(5, '0');

        this.opCode = this.initOpCode(instructionString);
        this.parameterModes = this.initParameterModes(instructionString);
    }

    public isDone(): boolean {
        return this.opCode === OpCode.Done;
    }

    public getOpCode(): number {
        return this.opCode;
    }

    public getParameterModes(): number[] {
        return this.parameterModes;
    }

    private initOpCode(instructionString: string): number {
        const opCodeString = instructionString.substring(instructionString.length - 2);

        return Number.parseInt(opCodeString, 10);
    }

    private initParameterModes(instructionString: string): number[] {
        if (instructionString.length <= 2) {
            return [];
        }

        return instructionString
            .substring(0, instructionString.length - 2)
            .split('')
            .map((parameterMode: string) => Number.parseInt(parameterMode))
            .reverse();
    }

}

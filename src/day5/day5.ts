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
    test1();
    test2();
    puzzle();
}

function test1(): void {
    const intCode = executeIntCode([1002, 4, 3, 4, 33]);

    assert.strictEqual(intCode[4], 99);
}

function test2(): void {
    const intCode = executeIntCode([1101, 100, -1, 4, 0]);

    assert.strictEqual(intCode[4], 99);
}

function puzzle(): void {
    executeIntCode(day5Input);
}

function executeIntCode(intCode: number[]): number[] {
    let instructionPointer: number = 0;

    while (true) {
        const instructionValue = intCode[instructionPointer];
        const instruction = new Instruction(instructionValue);

        if (instruction.isDone()) {
            return intCode;
        }

        const opCode = instruction.getOpCode();
        const parameterModes = instruction.getParameterModes();
        const pointerShift = executeOperation(intCode, opCode, instructionPointer, parameterModes);
        instructionPointer += pointerShift;
    }
}

function executeOperation(intCode: number[], opCode: number, instructionPointer: number, parameterModes: number[]): number {
    if (opCode === OpCode.Addition) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left + right;

        return 4;
    }

    if (opCode === OpCode.Multiplication) {
        const left = getParameter(intCode, instructionPointer, parameterModes, 0);
        const right = getParameter(intCode, instructionPointer, parameterModes, 1);
        const outputPointer = intCode[instructionPointer + 3];
        intCode[outputPointer] = left * right;

        return 4;
    }

    if (opCode === OpCode.Input) {
        const outputPointer = intCode[instructionPointer + 1];
        const input = readLineSync.question(`Input for position ${outputPointer}: `);
        intCode[outputPointer] = Number.parseInt(input, 10);

        return 2;
    }

    if (opCode === OpCode.Output) {
        const outputPointer = intCode[instructionPointer + 1];
        const output = intCode[outputPointer];

        consola.info(`Value at position ${outputPointer} is ${output}`);

        return 2;
    }

    throw new Error(`Unknown opcode: ${opCode}!`);
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

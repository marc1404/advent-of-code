import { Instruction } from './Instruction';
import { OpCode } from './OpCode';
import consola from 'consola';
import readLineSync from 'readline-sync';
import { ParameterMode } from './ParameterMode';

export class IntCode {

    private readonly intCode: number[];
    private readonly inputs: number[];
    private readonly outputs: number[] = [];
    private instructionPointer: number = 0;
    private relativeBase: number = 0;
    private _isDone: boolean = false;

    constructor(intCode: number[], inputs: number[] = [], relativeBase: number = 0) {
        this.intCode = [...intCode];
        this.inputs = [...inputs];
        this.relativeBase = relativeBase;
    }

    public addInput(input: number): void {
        this.inputs.push(input);
    }

    public getOutputs(): number[] {
        return this.outputs;
    }

    public getIntCode(): number[] {
        return this.intCode;
    }

    public isDone(): boolean {
        return this._isDone;
    }

    public getRelativeBase(): number {
        return this.relativeBase;
    }

    public execute(yieldOnOutput: boolean = false): IntCode {
        while (true) {
            const {instructionPointer} = this;
            const instructionValue = this.getIntCodeAt(instructionPointer);
            const instruction = new Instruction(instructionValue);

            if (instruction.isDone()) {
                this._isDone = true;

                return this;
            }

            const outputCount = this.outputs.length;
            this.instructionPointer = this.executeInstruction(instruction);
            const hasNewOutput = this.outputs.length > outputCount;

            if (yieldOnOutput && hasNewOutput) {
                return this;
            }
        }
    }

    private executeInstruction(instruction: Instruction): number {
        const {intCode, instructionPointer} = this;
        const opCode = instruction.getOpCode();
        const parameterModes = instruction.getParameterModes();

        if (opCode === OpCode.Addition) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = this.getPosition(parameterModes, 2);
            intCode[outputPointer] = left + right;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Multiplication) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = this.getPosition(parameterModes, 2);
            intCode[outputPointer] = left * right;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Input) {
            const outputPointer = this.getPosition(parameterModes, 0);
            intCode[outputPointer] = this.readInput(outputPointer);

            return instructionPointer + 2;
        }

        if (opCode === OpCode.Output) {
            const position = this.getPosition(parameterModes, 0);
            const output = this.getIntCodeAt(position);

            this.outputs.unshift(output);
            consola.info(`Value at position ${position} is ${output}`);

            return instructionPointer + 2;
        }

        if (opCode === OpCode.JumpIfTrue) {
            const a = this.getParameter(parameterModes, 0);
            const b = this.getParameter(parameterModes, 1);

            return a !== 0
                ? b
                : instructionPointer + 3;
        }

        if (opCode === OpCode.JumpIfFalse) {
            const a = this.getParameter(parameterModes, 0);
            const b = this.getParameter(parameterModes, 1);

            return a === 0
                ? b
                : instructionPointer + 3;
        }

        if (opCode === OpCode.LessThan) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = this.getPosition(parameterModes, 2);
            intCode[outputPointer] = left < right ? 1 : 0;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Equals) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = this.getPosition(parameterModes, 2);
            intCode[outputPointer] = left === right ? 1 : 0;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.RelativeBase) {
            const modifier = this.getParameter(parameterModes, 0);
            this.relativeBase += modifier;

            consola.info(`Relative base modified by ${modifier} is now ${this.relativeBase}`);

            return instructionPointer + 2;
        }

        throw new Error(`Unknown opcode: ${opCode}!`);
    }

    private readInput(outputPointer: number): number {
        const input = this.inputs.shift();

        if (input !== undefined) {
            return input;
        }

        const inputString = readLineSync.question(`Input for position ${outputPointer}: `);

        return Number.parseInt(inputString, 10);
    }

    private getParameter(parameterModes: number[], parameterIndex: number): number {
        const position = this.getPosition(parameterModes, parameterIndex);

        return this.getIntCodeAt(position);
    }

    private getPosition(parameterModes: number[], parameterIndex: number): number {
        const { instructionPointer } = this;
        const parameterMode = parameterModes[parameterIndex];
        const parameterPointer = instructionPointer + 1 + parameterIndex;

        if (parameterMode === ParameterMode.Position) {
            return this.getIntCodeAt(parameterPointer);
        }

        if (parameterMode === ParameterMode.Immediate) {
            return parameterPointer;
        }

        if (parameterMode === ParameterMode.Relative) {
            const parameter = this.getIntCodeAt(parameterPointer);

            return this.relativeBase + parameter;
        }

        throw new Error(`Unknown parameter mode: ${parameterMode}!`);
    }

    private getIntCodeAt(position: number): number {
        if (position < 0) {
            throw new Error(`Negative indices are not supported! (${position})`);
        }

        return this.intCode[position] ?? 0;
    }

}

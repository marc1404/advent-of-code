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
    private _isDone: boolean = false;

    constructor(intCode: number[], inputs: number[] = []) {
        this.intCode = [...intCode];
        this.inputs = [...inputs];
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

    public execute(yieldOnOutput: boolean = false): IntCode {
        while (true) {
            const {intCode, instructionPointer} = this;
            const instructionValue = intCode[instructionPointer];
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
            const outputPointer = intCode[instructionPointer + 3];
            intCode[outputPointer] = left + right;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Multiplication) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = intCode[instructionPointer + 3];
            intCode[outputPointer] = left * right;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Input) {
            const outputPointer = intCode[instructionPointer + 1];
            intCode[outputPointer] = this.readInput(outputPointer);

            return instructionPointer + 2;
        }

        if (opCode === OpCode.Output) {
            const outputPointer = intCode[instructionPointer + 1];
            const output = intCode[outputPointer];

            this.outputs.unshift(output);
            consola.info(`Value at position ${outputPointer} is ${output}`);

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
            const outputPointer = intCode[instructionPointer + 3];
            intCode[outputPointer] = left < right ? 1 : 0;

            return instructionPointer + 4;
        }

        if (opCode === OpCode.Equals) {
            const left = this.getParameter(parameterModes, 0);
            const right = this.getParameter(parameterModes, 1);
            const outputPointer = intCode[instructionPointer + 3];
            intCode[outputPointer] = left === right ? 1 : 0;

            return instructionPointer + 4;
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
        const {intCode, instructionPointer} = this;
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

}

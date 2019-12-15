import { OpCode } from './OpCode';

export class Instruction {

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

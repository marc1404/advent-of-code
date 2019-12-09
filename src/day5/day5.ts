import { getOperation, OpCode } from '../day2/day2';

export function day5(): void {
    test1();
}

function test1() {
    const intCode = [1002, 4, 3, 4, 33];
}

function executeIntCode(intCode: number[]): number[] {
    let instructionPointer: number = 0;

    while (true) {
        const instructionValue = intCode[instructionPointer];
        const instruction = new Instruction(instructionValue);

        if (instruction.isDone()) {
            return intCode;
        }

        const operation = instruction.getOperation();

        instructionPointer += instruction.getParameterCount() + 1;
    }
}

class Instruction {

    private readonly opCode: number;
    private readonly parameterModes: number[];

    constructor(
        private readonly instructionValue: number
    ) {
        const instructionString = instructionValue.toString();
        this.opCode = this.initOpCode(instructionString);
        this.parameterModes = this.initParameterModes(instructionString);
    }

    public isDone(): boolean {
        return this.opCode === OpCode.Done;
    }

    public getParameterCount(): number {
        return this.parameterModes.length;
    }

    public getOperation(): (left: number, right: number) => number {
        return getOperation(this.opCode);
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

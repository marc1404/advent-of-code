import { executeIntCode } from '../day5/day5';
import assert from 'assert';
import { day7Input } from './input';
import consola from 'consola';

export function day7(): void {
    test1();
    test2();
    test3();
    puzzle1();
}

function test1(): void {
    const intCode = [3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0];
    const phaseSettings = [4, 3, 2, 1, 0];
    const outputSignal = runAmplifiers(intCode, phaseSettings);

    assert.strictEqual(outputSignal, 43210);
}

function test2(): void {
    const intCode = [3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0];
    const phaseSettings = [0, 1, 2, 3, 4];
    const outputSignal = runAmplifiers(intCode, phaseSettings);

    assert.strictEqual(outputSignal, 54321);
}

function test3(): void {
    const intCode = [3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0];
    const phaseSettings = [1, 0, 4, 3, 2];
    const outputSignal = runAmplifiers(intCode, phaseSettings);

    assert.strictEqual(outputSignal, 65210);
}

function puzzle1(): void {
    const phaseSettings = [0, 1, 2, 3, 4];
    const maxOutputSignal = getPermutations(phaseSettings)
        .map(phaseSettings => runAmplifiers(day7Input, phaseSettings))
        .reduce((maxOutputSignal, outputSignal) => Math.max(outputSignal, maxOutputSignal), Number.MIN_VALUE);

    consola.info(`The highest output signal is ${maxOutputSignal}`);
}

function getPermutations(sequence: number[]): number[][] {
    const permutations: number[][] = [];

    if (sequence.length === 1) {
        permutations.push(sequence);

        return permutations;
    }

    for (let i = 0; i < sequence.length; i++) {
        const number = sequence[i];
        const remainingSequence = [
            ...sequence.slice(0, i),
            ...sequence.slice(i + 1)
        ];

        const innerPermutations = getPermutations(remainingSequence)
            .map(innerPermutation => [number, ...innerPermutation]);

        permutations.push(...innerPermutations);
    }

    return permutations;
}

function runAmplifiers(intCode: number[], phaseSettings: number[]): number {
    let lastOutputSignal = 0;

    for (const phaseSetting of phaseSettings) {
        const amplifierIntCode = [...intCode];
        const inputs = [phaseSetting, lastOutputSignal];
        const outputs: number[] = [];

        executeIntCode(amplifierIntCode, inputs, outputs);

        lastOutputSignal = outputs[0];
    }

    return lastOutputSignal;
}

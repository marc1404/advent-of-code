import { IntCode } from '../int-code/IntCode';
import assert from 'assert';
import consola from 'consola';
import { day9Input } from './input';

export function day9(): void {
    test1();
    test2();
    test3();
    test4();
    test5();
    puzzle1();
    puzzle2();
}

function test1(): void {
    consola.start('Test 1');

    const intCode = new IntCode([109, 19, 99], [], 2000);

    intCode.execute();

    assert.strictEqual(intCode.getRelativeBase(), 2019);
    consola.success('Test 1');
}

function test2(): void {
    consola.start('Test 2');

    const outputs = new IntCode([109, 19, 204, -34, 99], [], 2000)
        .execute()
        .getOutputs();

    assert.deepStrictEqual(outputs, [0]);
    consola.success('Test 2');
}

function test3(): void {
    consola.start('Test 3');

    const intCode = [109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99];
    const outputs = new IntCode(intCode)
        .execute()
        .getOutputs()
        .reverse();

    assert.deepStrictEqual(outputs, intCode);
    consola.success('Test 3');
}

function test4(): void {
    consola.start('Test 4');
    consola.info('Should output a 16-digit number');

    const outputs = new IntCode([1102, 34915192, 34915192, 7, 4, 7, 99, 0])
        .execute()
        .getOutputs();

    assert.deepStrictEqual(outputs, [34915192 * 34915192]);
    consola.success('Test 4');
}

function test5(): void {
    consola.start('Test 5');

    const outputs = new IntCode([104, 1125899906842624, 99])
        .execute()
        .getOutputs();

    assert.deepStrictEqual(outputs, [1125899906842624]);
    consola.success('Test 5');
}

function puzzle1(): void {
    consola.start('Puzzle 1');

    const outputs = new IntCode(day9Input, [1])
        .execute()
        .getOutputs()
        .reverse();

    consola.info('Outputs:', outputs);
    consola.success('Puzzle 1');
}

function puzzle2(): void {
    consola.start('Puzzle 2');

    const outputs = new IntCode(day9Input, [2])
        .execute()
        .getOutputs()
        .reverse();

    consola.info('Outputs:', outputs);
    consola.success('Puzzle 2');
}

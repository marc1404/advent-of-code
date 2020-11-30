import assert from 'assert';
import consola from 'consola';

const lowerBound = 264793;
const upperBound = 803935;

export function day4(): void {
    test1();
    test2();
    test3();
    test4();
    test5();
    test6();
    puzzle();
}

function test1(): void {
    const actual = isValidPassword(111111, {ignoreRange: true});
    const expected = true;

    assert.strictEqual(actual, expected);
}

function test2(): void {
    const actual = isValidPassword(223450, {ignoreRange: true});
    const expected = false;

    assert.strictEqual(actual, expected);
}

function test3(): void {
    const actual = isValidPassword(123789, {ignoreRange: true});
    const expected = false;

    assert.strictEqual(actual, expected);
}

function test4(): void {
    const actual = isValidPassword(112233, {ignoreRange: true, strictTwoSameAdjacentDigits: true});
    const expected = true;

    assert.strictEqual(actual, expected);
}

function test5(): void {
    const actual = isValidPassword(123444, {ignoreRange: true, strictTwoSameAdjacentDigits: true});
    const expected = false;

    assert.strictEqual(actual, expected);
}

function test6(): void {
    const actual = isValidPassword(111122, {ignoreRange: true, strictTwoSameAdjacentDigits: true});
    const expected = true;

    assert.strictEqual(actual, expected);
}

function puzzle(): void {
    const passwords = [];

    for (let password = lowerBound; password <= upperBound; password++) {
        passwords.push(password);
    }

    const validPasswords = passwords.filter(password => isValidPassword(password));
    const strictValidPasswords = validPasswords.filter(password => isValidPassword(password, {strictTwoSameAdjacentDigits: true}));

    consola.info(`There are ${validPasswords.length} valid passwords`);
    consola.info(`There are ${strictValidPasswords.length} strict valid passwords`);
}

interface passwordValidationOptions {
    ignoreRange?: boolean;
    strictTwoSameAdjacentDigits?: boolean;
}

function isValidPassword(password: number, options?: passwordValidationOptions): boolean {
    const sequence = password.toString();
    const isSixDigits = sequence.length === 6;
    const isWithingRange = (password >= lowerBound && password <= upperBound) || !!options?.ignoreRange;
    const hasTwoSameAdjacentDigits = checkForTwoSameAdjacentDigits(sequence, !!options?.strictTwoSameAdjacentDigits);
    const hasOnlyIncreasingDigits = checkForOnlyIncreasingDigits(sequence);

    return isSixDigits && isWithingRange && hasTwoSameAdjacentDigits && hasOnlyIncreasingDigits;
}

function checkForTwoSameAdjacentDigits(sequence: string, strict: boolean = false): boolean {
    for (let i = 0; i < sequence.length - 1; i++) {
        const z = sequence[i - 1];
        const a = sequence[i];
        const b = sequence[i + 1];
        const c = sequence[i + 2];
        const zDoesNotEqualA = z !== a;
        const aEqualsB = a === b;
        const bDoesNotEqualC = b !== c;

        if (!strict && aEqualsB) {
            return true;
        }

        if (strict && zDoesNotEqualA && aEqualsB && bDoesNotEqualC) {
            return true;
        }
    }

    return false;
}

function checkForOnlyIncreasingDigits(sequence: string): boolean {
    const digits = sequence
        .split('')
        .map((char: string) => Number.parseInt(char, 10));

    for (let i = 0; i < digits.length - 1; i++) {
        if (digits[i] > digits[i + 1]) {
            return false;
        }
    }

    return true;
}

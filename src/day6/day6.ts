import fs from 'fs';
import path from 'path';
import os from 'os';
import consola from 'consola';

export function day6(): void {
    consola.info(readInputLines());
}

function readInputLines(): string[] {
    const inputPath = path.join(__dirname, 'input.txt');
    const input = fs.readFileSync(inputPath, {encoding: 'utf8'});

    return input.split(os.EOL);
}

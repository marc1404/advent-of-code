import fs from 'fs';
import path from 'path';
import os from 'os';
import consola from 'consola';

export function day1(): void {
    const inputPath = path.join(__dirname, 'input.txt');
    const input = fs.readFileSync(inputPath, {encoding: 'utf8'});
    const lines = input.split(os.EOL);

    const moduleMasses: number[] = lines
        .map(line => Number.parseInt(line, 10))
        .filter(moduleMass => !Number.isNaN(moduleMass));

    const totalFuel = moduleMasses.reduce((totalFuel, moduleMass) => {
        return totalFuel + calculateFuel(moduleMass);
    }, 0);

    consola.info(`Total fuel requirement: ${totalFuel}`);

    const totalFuelWithCompensation = moduleMasses.reduce((totalFuel, moduleMass) => {
        return totalFuel + calculateFuel(moduleMass, true);
    }, 0);

    consola.info(`Total fuel requirement with compoensation: ${totalFuelWithCompensation}`);
}

function calculateFuel(mass: number, compensateFuel: boolean = false): number {
    let fuel = Math.floor(mass / 3) - 2;
    fuel = Math.max(fuel, 0);

    if (fuel === 0) {
        return fuel;
    }

    if (compensateFuel) {
        fuel += calculateFuel(fuel, true);
    }

    return fuel;
}

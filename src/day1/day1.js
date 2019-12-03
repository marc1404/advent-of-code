const fs = require('fs');
const path = require('path');
const os = require('os');
const consola = require('consola');

const inputPath = path.join(__dirname, 'input.txt');
const input = fs.readFileSync(inputPath, { encoding: 'utf8' });
const lines = input.split(os.EOL);

const moduleMasses = lines.map(line => Number.parseInt(line, 10));

const totalFuel = moduleMasses.reduce((totalFuel, moduleMass) => {
    if (Number.isNaN(moduleMass)) {
        return totalFuel;
    }

    return totalFuel + Math.floor(moduleMass / 3) - 2;
}, 0);

consola.success(`Total fuel requirement: ${totalFuel}`);

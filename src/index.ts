import consola from 'consola';
import {day1} from './day1/day1';

const dayToFunction: Record<string, () => void> = {
    1: day1
};

const day = process.argv[2];

if (!day) {
    consola.error('Please specify the day that you would like to execute e.g: yarn start 1');
    process.exit(1);
}

const dayFunction = dayToFunction[day];

if (!dayFunction) {
    consola.error(`Unknown day: ${day}!`);
    process.exit(1);
}

consola.start(`Executing day ${day}`);

dayFunction();

consola.success('Done!');

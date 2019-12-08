import consola from 'consola';
import { day1 } from './day1/day1';
import { day2 } from './day2/day2';
import { day3 } from './day3/day3';
import { day4 } from './day4/day4';
import { day5 } from './day5/day5';

const dayToFunction: Record<string, () => void> = {
    1: day1,
    2: day2,
    3: day3,
    4: day4,
    5: day5
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

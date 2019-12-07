import consola from 'consola';
import assert from 'assert';
import fs from 'fs';
import path from 'path';
import os from 'os';

export function day3(): void {
    test1();
    test2();
    test3();
    puzzle();
}

function test1(): void {
    const wire1 = new Wire(['R8', 'U5', 'L5', 'D3']);
    const wire2 = new Wire(['U7', 'R6', 'D4', 'L4']);
    const crossings = determineCrossings(wire1, wire2);
    const distance = determineDistanceOfClosestCrossing(crossings);
    const totalSteps = determineLeastTotalSteps(crossings);

    assert.strictEqual(distance, 6);
    assert.strictEqual(totalSteps, 30);
}

function test2(): void {
    const wire1 = new Wire(['R75', 'D30', 'R83', 'U83', 'L12', 'D49', 'R71', 'U7', 'L72']);
    const wire2 = new Wire(['U62', 'R66', 'U55', 'R34', 'D71', 'R55', 'D58', 'R83']);
    const crossings = determineCrossings(wire1, wire2);
    const distance = determineDistanceOfClosestCrossing(crossings);
    const totalSteps = determineLeastTotalSteps(crossings);

    assert.strictEqual(distance, 159);
    assert.strictEqual(totalSteps, 610);
}

function test3(): void {
    const wire1 = new Wire(['R98', 'U47', 'R26', 'D63', 'R33', 'U87', 'L62', 'D20', 'R33', 'U53', 'R51']);
    const wire2 = new Wire(['U98', 'R91', 'D20', 'R16', 'D67', 'R40', 'U7', 'R15', 'U6', 'R7']);
    const crossings = determineCrossings(wire1, wire2);
    const distance = determineDistanceOfClosestCrossing(crossings);
    const totalSteps = determineLeastTotalSteps(crossings);

    assert.strictEqual(distance, 135);
    assert.strictEqual(totalSteps, 410);
}

function puzzle(): void {
    const inputPath = path.join(__dirname, 'input.txt');
    const input = fs.readFileSync(inputPath, {encoding: 'utf8'});
    const [line1, line2] = input.split(os.EOL);
    const wire1Path = line1.split(',');
    const wire2Path = line2.split(',');
    const wire1 = new Wire(wire1Path);
    const wire2 = new Wire(wire2Path);
    const crossings = determineCrossings(wire1, wire2);
    const distance = determineDistanceOfClosestCrossing(crossings);
    const totalSteps = determineLeastTotalSteps(crossings);

    consola.info(`Distance of closest crossing: ${distance}`);
    consola.info(`Fewest combined steps to reach an intersection: ${totalSteps}`);
}

function determineDistanceOfClosestCrossing(crossings: Crossing[]): number {
    const [closestDistance] = crossings
        .map(crossing => crossing.getManhattanDistanceToOrigin())
        .sort((a, b) => a - b);

    return closestDistance;
}

function determineLeastTotalSteps(crossings: Crossing[]): number {
    return crossings
        .map(crossing => crossing.getTotalSteps())
        .reduce((leastTotalSteps: number, totalSteps: number) => {
            return totalSteps < leastTotalSteps ? totalSteps : leastTotalSteps;
        }, Number.MAX_VALUE);
}

function determineCrossings(wire1: Wire, wire2: Wire): Crossing[] {
    const panel: Map<string, Coordinate> = new Map<string, Coordinate>();
    const coordinates1 = wire1.getCoordinates();
    const coordinates2 = wire2.getCoordinates();

    coordinates1.forEach(coordinate => panel.set(coordinate.toString(), coordinate));

    return coordinates2
        .filter(coordinate => panel.has(coordinate.toString()))
        .map(coordinate2 => {
            const coordinate1 = panel.get(coordinate2.toString());

            return new Crossing(coordinate1 as Coordinate, coordinate2);
        })
        .filter(crossing => !!crossing);
}

enum Directions {
    Up = 'U',
    Right = 'R',
    Down = 'D',
    Left = 'L'
}

class Wire {

    private readonly coordinates: Coordinate[];

    constructor(
        private readonly path: string[]
    ) {
        this.coordinates = this.initCoordinates();
    }

    public getCoordinates(): Coordinate[] {
        return this.coordinates;
    }

    private initCoordinates(): Coordinate[] {
        const coordinates: Coordinate[] = [];
        let x = 0;
        let y = 0;
        let steps = 0;

        for (const directionInput of this.path) {
            const direction = new Direction(directionInput);
            const [modX, modY] = direction.getCoordinateModifiers();

            for (let i = 0; i < direction.getSteps(); i++) {
                x += modX;
                y += modY;
                steps++;

                const coordinate = new Coordinate(x, y, steps);

                coordinates.push(coordinate);
            }
        }

        return coordinates;
    }

}

class Direction {

    private readonly coordinateModifiers: Record<string, [number, number]> = {
        [Directions.Up]: [0, -1],
        [Directions.Right]: [1, 0],
        [Directions.Down]: [0, 1],
        [Directions.Left]: [-1, 0]
    };

    private readonly direction: string;
    private readonly steps: number;

    constructor(directionInput: string) {
        this.direction = directionInput[0];
        this.steps = Number.parseInt(directionInput.substring(1));
    }

    getCoordinateModifiers(): [number, number] {
        const coordinateModifier = this.coordinateModifiers[this.direction];

        if (!coordinateModifier) {
            throw new Error(`Unknown direction: ${this.direction}!`);
        }

        return coordinateModifier;
    }

    getSteps(): number {
        return this.steps;
    }

}

class Coordinate {

    constructor(
        private readonly x: number,
        private readonly y: number,
        private readonly steps: number
    ) {
    }

    getManhattanDistanceToOrigin(): number {
        return Math.abs(this.x) + Math.abs(this.y);
    }

    getSteps(): number {
        return this.steps;
    }

    toString(): string {
        return `${this.x}:${this.y}`;
    }

}

class Crossing {

    private readonly totalSteps: number;

    constructor(
        private readonly coordinate1: Coordinate,
        private readonly coordinate2: Coordinate
    ) {
        this.totalSteps = coordinate1.getSteps() + coordinate2.getSteps();
    }

    getManhattanDistanceToOrigin(): number {
        return this.coordinate1.getManhattanDistanceToOrigin();
    }

    getTotalSteps(): number {
        return this.totalSteps;
    }

}

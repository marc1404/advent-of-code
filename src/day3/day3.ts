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

    function test1(): void {
        const wire1 = new Wire(['R8', 'U5', 'L5', 'D3']);
        const wire2 = new Wire(['U7', 'R6', 'D4', 'L4']);
        const distance = determineDistanceOfClosestCrossing(wire1, wire2);

        assert.strictEqual(distance, 6);
    }

    function test2(): void {
        const wire1 = new Wire(['R75', 'D30', 'R83', 'U83', 'L12', 'D49', 'R71', 'U7', 'L72']);
        const wire2 = new Wire(['U62', 'R66', 'U55', 'R34', 'D71', 'R55', 'D58', 'R83']);
        const distance = determineDistanceOfClosestCrossing(wire1, wire2);

        assert.strictEqual(distance, 159);
    }

    function test3(): void {
        const wire1 = new Wire(['R98', 'U47', 'R26', 'D63', 'R33', 'U87', 'L62', 'D20', 'R33', 'U53', 'R51']);
        const wire2 = new Wire(['U98', 'R91', 'D20', 'R16', 'D67', 'R40', 'U7', 'R15', 'U6', 'R7']);
        const distance = determineDistanceOfClosestCrossing(wire1, wire2);

        assert.strictEqual(distance, 135);
    }

    function puzzle(): void {
        const inputPath = path.join(__dirname, 'input.txt');
        const input = fs.readFileSync(inputPath, {encoding: 'utf8'});
        const [line1, line2] = input.split(os.EOL);
        const wire1Path = line1.split(',');
        const wire2Path = line2.split(',');
        const wire1 = new Wire(wire1Path);
        const wire2 = new Wire(wire2Path);
        const distance = determineDistanceOfClosestCrossing(wire1, wire2);

        consola.info(`Distance of closest crossing: ${distance}`);
    }
}

function determineDistanceOfClosestCrossing(wire1: Wire, wire2: Wire): number {
    const panel: Set<string> = new Set<string>();
    const coordinates1 = wire1.getCoordinates();
    const coordinates2 = wire2.getCoordinates();

    coordinates1.forEach(coordinate => panel.add(coordinate.toString()));

    const crossings = coordinates2.filter(coordinate => panel.has(coordinate.toString()));
    const [closestDistance] = crossings
        .map(crossing => crossing.getManhattanDistanceToOrigin())
        .sort((a, b) => a - b);

    return closestDistance;
}

enum Directions {
    Up = 'U',
    Right = 'R',
    Down = 'D',
    Left = 'L'
}

class Wire {

    constructor(
        private readonly path: string[]
    ) {
    }

    getCoordinates(): Coordinate[] {
        const coordinates: Coordinate[] = [];
        let x = 0;
        let y = 0;

        for (const directionInput of this.path) {
            const direction = new Direction(directionInput);
            const [modX, modY] = direction.getCoordinateModifiers();

            for (let i = 0; i < direction.getSteps(); i++) {
                x += modX;
                y += modY;

                coordinates.push(new Coordinate(x, y));
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
        private readonly y: number
    ) {
    }

    getManhattanDistanceToOrigin(): number {
        return Math.abs(this.x) + Math.abs(this.y);
    }

    toString(): string {
        return `${this.x}:${this.y}`;
    }

}

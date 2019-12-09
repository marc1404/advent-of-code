import fs from 'fs';
import path from 'path';
import os from 'os';
import consola from 'consola';
import assert from 'assert';

export function day6(): void {
    test();
    puzzle1();
}

function test(): void {
    const orbits = [
        'COM)B',
        'B)C',
        'C)D',
        'D)E',
        'E)F',
        'B)G',
        'G)H',
        'D)I',
        'E)J',
        'J)K',
        'K)L'
    ];

    const checksum = calculateOrbitCountChecksum(orbits);

    assert.strictEqual(checksum, 42);
}

function puzzle1(): void {
    const orbits = readInputLines();
    const orbitCountChecksum = calculateOrbitCountChecksum(orbits);

    consola.info(`The orbit count checksum is ${orbitCountChecksum}`);
}

function readInputLines(): string[] {
    const inputPath = path.join(__dirname, 'input.txt');
    const input = fs.readFileSync(inputPath, {encoding: 'utf8'});

    return input
        .split(os.EOL)
        .filter(line => !!line);
}

function calculateOrbitCountChecksum(orbits: string[]): number {
    const identifierToSpaceObject: Map<string, SpaceObject> = new Map<string, SpaceObject>();

    orbits.forEach(orbit => parseOrbit(identifierToSpaceObject, orbit));

    const spaceObjects = Array.from(identifierToSpaceObject.values());

    return spaceObjects.reduce((orbitCountChecksum: number, spaceObject: SpaceObject) => orbitCountChecksum + spaceObject.countOrbits(), 0);
}

function parseOrbit(identifierToSpaceObject: Map<string, SpaceObject>, orbit: string): void {
    if (!orbit) {
        return;
    }

    const [leftIdentifier, rightIdentifier] = orbit.split(')');

    const leftObject = getSpaceObject(identifierToSpaceObject, leftIdentifier);
    const rightObject = getSpaceObject(identifierToSpaceObject, rightIdentifier);

    leftObject.addChild(rightObject);
    rightObject.setParent(leftObject);
}

function getSpaceObject(identifierToSpaceObject: Map<string, SpaceObject>, identifier: string): SpaceObject {
    const spaceObject = identifierToSpaceObject.get(identifier);

    if (spaceObject) {
        return spaceObject;
    }

    const newSpaceObject = new SpaceObject(identifier);

    identifierToSpaceObject.set(identifier, newSpaceObject);

    return newSpaceObject;
}

class SpaceObject {

    private readonly children: SpaceObject[] = [];
    private parent: SpaceObject | null = null;

    constructor(
        private readonly identifier: string
    ) {
    }

    public addChild(spaceObject: SpaceObject): void {
        this.children.push(spaceObject);
    }

    public setParent(spaceObject: SpaceObject): void {
        this.parent = spaceObject;
    }

    public countOrbits(): number {
        const {parent} = this;

        if (!parent) {
            return 0;
        }

        return 1 + parent.countOrbits();
    }

}

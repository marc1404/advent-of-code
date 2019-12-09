import fs from 'fs';
import path from 'path';
import os from 'os';
import consola from 'consola';
import assert from 'assert';

export function day6(): void {
    const puzzleOrbits = readInputLines();
    const puzzleOrbitMap = createOrbitMap(puzzleOrbits);

    test1();
    puzzle1(puzzleOrbits);
    test2();
    puzzle2(puzzleOrbitMap);
}

function test1(): void {
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

function puzzle1(orbits: string[]): void {
    const orbitCountChecksum = calculateOrbitCountChecksum(orbits);

    consola.info(`The orbit count checksum is ${orbitCountChecksum}`);
}

function test2(): void {
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
        'K)L',
        'K)YOU',
        'I)SAN'
    ];

    const orbitMap = createOrbitMap(orbits);
    const youObject = orbitMap.get('YOU') as SpaceObject;
    const sanObject = orbitMap.get('SAN') as SpaceObject;

    assert.notStrictEqual(youObject, undefined);
    assert.notStrictEqual(sanObject, undefined);

    const orbitalTransfers = calculateOrbitalTransfers(youObject, sanObject);

    assert.strictEqual(orbitalTransfers, 4);
}

function puzzle2(orbitMap: Map<string, SpaceObject>): void {
    const youObject = orbitMap.get('YOU') as SpaceObject;
    const sanObject = orbitMap.get('SAN') as SpaceObject;

    assert.notStrictEqual(youObject, undefined);
    assert.notStrictEqual(sanObject, undefined);

    const orbitalTransfers = calculateOrbitalTransfers(youObject, sanObject);

    consola.info(`The minimum orbital transfers required to move from YOU to SAN are ${orbitalTransfers}`);
}

function readInputLines(): string[] {
    const inputPath = path.join(__dirname, 'input.txt');
    const input = fs.readFileSync(inputPath, {encoding: 'utf8'});

    return input
        .split(os.EOL)
        .filter(line => !!line);
}

function calculateOrbitalTransfers(leftObject: SpaceObject, rightObject: SpaceObject): number {
    const leftPath = leftObject.getPathToRootObject().map(spaceObject => spaceObject.getIdentifier());
    const rightPath = rightObject.getPathToRootObject().map(spaceObject => spaceObject.getIdentifier());
    const connectingIdentifier = leftPath.find(identifier => rightPath.includes(identifier));
    const leftOrbitalTransfers = calculateOrbitalTransfersUntil(leftPath, connectingIdentifier as string);
    const rightOrbitalTransfers = calculateOrbitalTransfersUntil(rightPath, connectingIdentifier as string);

    return leftOrbitalTransfers + rightOrbitalTransfers;
}

function calculateOrbitalTransfersUntil(identifiers: string[], connectingIdentifier: string): number {
    for (let i = 0; i < identifiers.length; i++) {
        if (identifiers[i] === connectingIdentifier) {
            return i - 1;
        }
    }

    return 0;
}

function calculateOrbitCountChecksum(orbits: string[]): number {
    const orbitMap: Map<string, SpaceObject> = createOrbitMap(orbits);

    const spaceObjects = Array.from(orbitMap.values());

    return spaceObjects.reduce((orbitCountChecksum: number, spaceObject: SpaceObject) => orbitCountChecksum + spaceObject.countOrbits(), 0);
}

function createOrbitMap(orbits: string[]): Map<string, SpaceObject> {
    const orbitMap: Map<string, SpaceObject> = new Map<string, SpaceObject>();

    orbits.forEach(orbit => parseOrbit(orbitMap, orbit));

    return orbitMap;
}

function parseOrbit(orbitMap: Map<string, SpaceObject>, orbit: string): void {
    if (!orbit) {
        return;
    }

    const [leftIdentifier, rightIdentifier] = orbit.split(')');

    const leftObject = getSpaceObject(orbitMap, leftIdentifier);
    const rightObject = getSpaceObject(orbitMap, rightIdentifier);

    leftObject.addChild(rightObject);
    rightObject.setParent(leftObject);
}

function getSpaceObject(orbitMap: Map<string, SpaceObject>, identifier: string): SpaceObject {
    const spaceObject = orbitMap.get(identifier);

    if (spaceObject) {
        return spaceObject;
    }

    const newSpaceObject = new SpaceObject(identifier);

    orbitMap.set(identifier, newSpaceObject);

    return newSpaceObject;
}

class SpaceObject {

    private readonly children: SpaceObject[] = [];
    private parent: SpaceObject | null = null;

    constructor(
        private readonly identifier: string
    ) {
    }

    public getIdentifier(): string {
        return this.identifier;
    }

    public addChild(spaceObject: SpaceObject): void {
        this.children.push(spaceObject);
    }

    public setParent(spaceObject: SpaceObject): void {
        this.parent = spaceObject;
    }

    public countOrbits(): number {
        return this.getPathToRootObject().length;
    }

    public getPathToRootObject(): SpaceObject[] {
        const {parent} = this;

        if (!parent) {
            return [];
        }

        return [
            this,
            ...parent.getPathToRootObject()
        ];
    }

}

import assert from 'assert';
import {day8Input} from './input';
import consola from 'consola';

enum Color {
    Black = 0,
    White = 1,
    Transparent = 2
}

export function day8(): void {
    test1();
    puzzle1();
    test2();
    puzzle2();
}

function test1(): void {
    const layers = getImageLayers('123456789012', 3, 2);

    assert.strictEqual(layers.length, 2);
}

function puzzle1(): void {
    const layers = getImageLayers(day8Input, 25, 6);

    assert.strictEqual(layers.length, day8Input.length / (25 * 6));

    const sortedLayers = layers.sort((a, b) => a.getDigitCount(0) - b.getDigitCount(0));
    const [layer] = sortedLayers;
    const output = layer.getDigitCount(1) * layer.getDigitCount(2);

    consola.info(output);
}

function test2(): void {
    const layers = getImageLayers('0222112222120000', 2, 2);
    const decodedImage = decodeImage(layers);

    assert.deepStrictEqual(decodedImage, [0, 1, 1, 0]);
    printImage(decodedImage, 2);
}

function puzzle2(): void {
    const layers = getImageLayers(day8Input, 25, 6);
    const decodedImage = decodeImage(layers);

    printImage(decodedImage, 25);
}

function decodeImage(layers: Layer[]): number[] {
    const decodedImage: number[] = [];
    const [layer] = layers;
    const length = layer.getLength();

    for (let i = 0; i < length; i++) {
        const pixel = determinePixelColor(layers, i);

        decodedImage.push(pixel);
    }

    return decodedImage;
}

function printImage(pixels: number[], width: number): void {
    const image = pixels
        .join('');

    for (let i = 0; i < image.length; i += width) {
        const row = image
            .substring(i, i + width)
            .replace(new RegExp(Color.Black.toString(), 'g'), '⬛️')
            .replace(new RegExp(Color.White.toString(), 'g'), '⬜️');

        consola.info(row);
    }
}

function determinePixelColor(layers: Layer[], pixel: number): number {
    for (const layer of layers) {
        const color = layer.getColorAt(pixel);

        if (color !== Color.Transparent) {
            return color;
        }
    }

    throw new Error(`Could not determine color for pixel ${pixel}!`);
}

function getImageLayers(imageData: string, width: number, height: number): Layer[] {
    const layerLength = width * height;
    const layers: Layer[] = [];

    for (let i = 0; i < imageData.length; i += layerLength) {
        const layerString = imageData.substring(i, i + layerLength);
        const layer = new Layer(layerString);

        layers.push(layer);
    }

    return layers;
}

class Layer {

    private readonly counts: number[];

    constructor(
        private readonly layerString: string
    ) {
        this.counts = [
            this.count('0'),
            this.count('1'),
            this.count('2')
        ];
    }

    public getLength(): number {
        return this.layerString.length;
    }

    public getDigitCount(digit: number): number {
        return this.counts[digit];
    }

    public getColorAt(pixel: number): number {
        return Number.parseInt(this.layerString[pixel], 10);
    }

    private count(digit: string): number {
        return this.layerString.match(new RegExp(digit, 'g'))?.length ?? 0;
    }

}

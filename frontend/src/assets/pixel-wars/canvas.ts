import type CanvasElementController from './CanvasController';
import type { Color } from './CanvasController';

export const ColorPallete: Color[] = [
  [0x00, 0x00, 0x00, 0xff],
  [0x00, 0x55, 0x00, 0xff],
  [0x00, 0xab, 0x00, 0xff],
  [0x00, 0xff, 0x00, 0xff],
  [0x00, 0x00, 0xff, 0xff],
  [0x00, 0x55, 0xff, 0xff],
  [0x00, 0xab, 0xff, 0xff],
  [0x00, 0xff, 0xff, 0xff],
  [0xff, 0x00, 0x00, 0xff],
  [0xff, 0x55, 0x00, 0xff],
  [0xff, 0xab, 0x00, 0xff],
  [0xff, 0xff, 0x00, 0xff],
  [0xff, 0x00, 0xff, 0xff],
  [0xff, 0x55, 0xff, 0xff],
  [0xff, 0xab, 0xff, 0xff],
  [0xff, 0xff, 0xff, 0xff]
];

export function encodeColor([r, g, b, a]: Color): number {
  return ColorPallete.indexOf([r, g, b, a]);
}

export function decodeColor(color: number): Color {
  return ColorPallete[color];
}

export async function initialLoad(canvasController: CanvasElementController) {
  let canvasResponse: Response;
  try {
    canvasResponse = await fetch('http://localhost:8080/api/canvas');
  } catch (err) {
    console.log(err);
  }

  const canvasJSON = await canvasResponse.json();
  const canvasString = canvasJSON['canvas'] as string;
  const canvasSize = canvasJSON['size'] as { width: number; height: number};

  
  const canvasPixels = [];
  for (let i = 0; i < canvasString.length; i++) {
    canvasPixels.push(canvasString.charCodeAt(i) & 15);
    canvasPixels.push(canvasString.charCodeAt(i) >> 4);
  }
  console.log(canvasPixels);

  const canvas = new Uint8ClampedArray(canvasSize.width * canvasSize.height * 4);
  for (let i = 0; i < canvasPixels.length; i++) {
    const offset = i * 4;
    const color: Color = decodeColor(canvasPixels[i]);
    canvas[offset] = color[0];
    canvas[offset + 1] = color[1];
    canvas[offset + 2] = color[2];
    canvas[offset + 3] = color[3];
  }

  const imageData = new ImageData(canvas, canvasSize.width, canvasSize.height);
  canvasController.ctx.putImageData(imageData, 0, 0);
}

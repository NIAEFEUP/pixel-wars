import type CanvasElementController from './CanvasController';

const ColorPallete = [
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

export async function initialLoad(canvasController: CanvasElementController) {
  let canvasResponse: Response;
  try {
    canvasResponse = await fetch('http://localhost:8080/api/canvas');
  } catch (err) {
    console.log(err);
  }

  const canvasPixels = [];

  const canvasVal = (await canvasResponse.json()) as string;

  for (let i = 0; i < canvasVal.length; i++) {
    canvasPixels.push(canvasVal.charCodeAt(i) & 15);
    canvasPixels.push(canvasVal.charCodeAt(i) >> 4);
  }
  console.log(canvasPixels);

  const canvas = new Uint8ClampedArray(252 * 252 * 4);
  for (let i = 0; i < canvasPixels.length; i++) {
    const offset = i * 4;
    const color = ColorPallete[canvasPixels[i]];
    canvas[offset] = color[0];
    canvas[offset + 1] = color[1];
    canvas[offset + 2] = color[2];
    canvas[offset + 3] = color[3];
  }

  const imageData = new ImageData(canvas, 252, 252);
  canvasController.ctx.putImageData(imageData, 0, 0);
}



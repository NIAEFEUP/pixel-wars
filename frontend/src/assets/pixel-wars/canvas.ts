import type CanvasElementController from './CanvasController';
import type { CanvasPixels, Color } from './CanvasController';

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
    canvasResponse = await fetch('pixelwars/api/canvas');
  } catch (err) {
    console.log(err);
  }
  const canvasJSON = await canvasResponse.json();
  const canvasString = atob(canvasJSON['canvas']);
  const bytes = new Uint8Array(canvasString.length);
  for(var i = 0; i < canvasString.length; i++){
    bytes[i] = canvasString.charCodeAt(i)
  }
  const canvasSize = canvasJSON['size'] as { width: number; height: number };

  const imageData = canvasStringToColorList(bytes, canvasSize);
  canvasController.putCanvasPixels(imageData);
}

function canvasStringToColorList(
  canvasArray: Uint8Array,
  canvasSize: { width: number; height: number }
): CanvasPixels {
  const canvasPixels = [];
  for (let i = 0; i < canvasArray.length; i++) {
    canvasPixels.push(canvasArray[i] >> 4);
    canvasPixels.push(canvasArray[i] & 15);
  }

  const colors: Color[] = [];
  for (let i = 0; i < canvasPixels.length; i++) {
    if(canvasPixels[i] > 15) console.log(canvasPixels[i]);
    let color: Color = decodeColor(canvasPixels[i]);
    colors.push(color);
  }

  return {
    colors,
    width:canvasSize.width,
    height:canvasSize.height
  };
}

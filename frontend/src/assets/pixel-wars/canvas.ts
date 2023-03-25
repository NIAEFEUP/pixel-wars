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
    canvasResponse = await fetch('http://localhost:8080/api/canvas');
  } catch (err) {
    console.log(err);
  }
  const canvasJSON = await canvasResponse.json();
  const canvasString = canvasJSON['canvas'] as string;
  const canvasSize = canvasJSON['size'] as { width: number; height: number };

  const imageData = canvasStringToColorList(canvasString, canvasSize);
  canvasController.putCanvasPixels(imageData);
}

function canvasStringToColorList(
  canvasString: string,
  canvasSize: { width: number; height: number }
): CanvasPixels {
  const canvasPixels = [];
  for (let i = 0; i < canvasString.length; i++) {
    canvasPixels.push(canvasString.charCodeAt(i) & 15);
    canvasPixels.push(canvasString.charCodeAt(i) >> 4);
  }

  const colors: Color[] = [];
  for (let i = 0; i < canvasPixels.length; i++) {
    
    let color: Color = decodeColor(canvasPixels[i]);
    if(color === undefined) color = ColorPallete[0];
    colors.push(color);
  }


  return {
    colors,
    width:canvasSize.width,
    height:canvasSize.height
  };
}

export default class CanvasElementController {
  canvas: HTMLCanvasElement;
  ctx: CanvasRenderingContext2D;

  constructor(canvas: HTMLCanvasElement) {
    this.canvas = canvas;
    this.canvas.width = window.innerWidth;
    this.canvas.height = window.innerHeight;
    this.setResizeListener();
    this.ctx = this.canvas.getContext('2d');
  }

  private setResizeListener() {
    window.addEventListener('resize', () => {
      this.canvas.width = window.innerWidth;
      this.canvas.height = window.innerHeight;

      this.ctx = this.canvas.getContext('2d');
    });
  }

  putPixel(x: number, y: number, [r, g, b, a]: Color) {
    this.ctx.fillStyle = `rgba(${r}, ${g}, ${b}, ${a})`;
    this.ctx.fillRect(x, y, 1, 1);
    this.ctx.fillStyle = undefined;
  }
}

export type Color = [number, number, number, number];

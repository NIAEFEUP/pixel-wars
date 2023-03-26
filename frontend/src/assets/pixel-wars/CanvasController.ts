export default class CanvasElementController {
  canvas: HTMLCanvasElement;
  ctx: CanvasRenderingContext2D;
  pixels: CanvasPixels;

  private scale = 1;
  private lastDir = false;
  private beginX = 0;
  private beginY = 0;
  private pressed = false;
  private moved = false;

  constructor(canvas: HTMLCanvasElement) {
    this.canvas = canvas;
    this.canvas.width = window.innerWidth;
    this.canvas.height = window.innerHeight;
    this.ctx = this.canvas.getContext('2d');
    this.ctx.imageSmoothingEnabled = false;
    this.setListeners();
  }

  private draw(transformCoords: { x: number, y: number }) {
    console.log(this.pixels);
    let matrix = this.ctx.getTransform();
    this.ctx.resetTransform();
    this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);
    matrix.e += transformCoords.x;
    matrix.f += transformCoords.y;
    this.ctx.setTransform(matrix);
    this.ctx.scale(this.scale, this.scale);
    for (let h = 0; h < this.pixels.height; h++) {
      for (let w = 0; w < this.pixels.width; w++) {
        this.putPixel(w, h, this.pixels.colors[w + (this.pixels.width * h)]);
      }
    }
  }

  private setListeners() {
    this.canvas.addEventListener('resize', () => {
      this.canvas.width = window.innerWidth;
      this.canvas.height = window.innerHeight;

      this.ctx = this.canvas.getContext('2d');
      this.draw({x:0,y:0});
    });
    this.canvas.addEventListener("wheel", (pog: WheelEvent) => {
      pog.preventDefault();
      if (pog.deltaY < 0 != this.lastDir) {
        this.lastDir = false;
        this.scale = 1;
      } else this.lastDir = true;
      this.scale += Math.min(pog.deltaY, 200) * -0.001;
      this.draw({ x: 0, y: 0 });
    }, { passive: false });

    this.canvas.addEventListener("mousedown", (pog: MouseEvent) => {
      pog.preventDefault();
      if (this.pressed == false && pog.button == 0) {
        this.scale = 1;
        this.beginX = pog.x;
        this.beginY = pog.y;
        this.pressed = true;
        return;
      }
    });
    this.canvas.addEventListener("mouseup", (pog) => {
      this.pressed = false;
      this.beginX = 0;
      this.beginY = 0;
      if(pog.button == 0 && this.moved==false) this.clickHandler(pog);
      this.moved = false;
      }
    );
    this.canvas.addEventListener("mousemove", (pog) => {
      if (this.pressed) {
        this.draw({ x: pog.x - this.beginX, y: pog.y - this.beginY });
        this.beginY = pog.y;
        this.beginX = pog.x;
        this.moved = true;
      }
    });

  }

  private clickHandler(mouseEvent : MouseEvent) {
    const canvasBoundingBox = this.canvas.getBoundingClientRect();
    const scale = this.ctx.getTransform().a;
    const xTransform = this.ctx.getTransform().e;
    const yTransform = this.ctx.getTransform().f;

    const pixelX = Math.floor((mouseEvent.x - canvasBoundingBox.x - xTransform) / scale);
    const pixelY = Math.floor((mouseEvent.y - canvasBoundingBox.y - yTransform) / scale);
    window.dispatchEvent(new CustomEvent("pixelClicked", {detail:{x:pixelX, y:pixelY}}));
    console.log(pixelX,pixelY);
  }

  putCanvasPixels(canvasPixels: CanvasPixels) {
    this.pixels = canvasPixels;
    this.scale = this.canvas.width / canvasPixels.width * 1.5;
    this.draw({ x: 0, y: 0 });
  }

  private putPixel(x: number, y: number, [r, g, b, a]: Color) {
    this.ctx.fillStyle = `rgba(${r}, ${g}, ${b}, ${a})`;
    this.ctx.fillRect(x, y, 1, 1);
  }

  putPixelCanvas(x: number, y: number, color: Color) {
    this.pixels.colors[x + (this.pixels.width * y)] = color;
    if(this.scale != 1) this.scale = 1;
    this.draw({ x: 0, y: 0 });
  }
}

export type Color = [number, number, number, number];

export type CanvasPixels = {
  colors: Color[]
  width: number
  height: number
}
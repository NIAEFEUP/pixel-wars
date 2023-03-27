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

  private zoomMobile = false;
  private prevDiff = 0;
  private prevPos = { x: 0, y: 0 };


  constructor(canvas: HTMLCanvasElement) {
    this.canvas = canvas;
    this.canvas.width = window.innerWidth;
    this.canvas.height = window.innerHeight;
    this.ctx = this.canvas.getContext('2d');
    this.ctx.imageSmoothingEnabled = false;
    this.setListeners();
  }

  private draw(transformCoords: { x: number, y: number }) {
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
      this.draw({ x: 0, y: 0 });
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
      if (pog.button == 0 && this.moved == false) this.clickHandler(pog);
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

    this.canvas.addEventListener("touchstart", (touch) => {
      touch.preventDefault();
      if (touch.touches.length == 2) {
        this.prevDiff = Math.hypot(touch.touches[0].clientX - touch.touches[1].clientX,
          touch.touches[0].clientY - touch.touches[1].clientY);
        this.zoomMobile = true;
        this.scale = 1;
      } else if (touch.touches.length == 1) {
        console.log(touch.touches);
        this.scale = 1;
        this.prevPos.x = touch.touches[0].clientX;
        this.prevPos.y = touch.touches[0].clientY;
        this.moved = true;
        this.pressed = true;
      }
    });
    this.canvas.addEventListener("touchmove", (touch) => {
      if (touch.touches.length == 2 && this.zoomMobile) {
        const diff = Math.hypot(touch.touches[0].clientX - touch.touches[1].clientX,
          touch.touches[0].clientY - touch.touches[1].clientY);
        if (diff - this.prevDiff == 0) return;
        this.scale += (diff - this.prevDiff) * 0.004;
        if (this.scale < 0.5) {
          this.scale = 0.5;
        } else if (this.scale > 1.5) {
          this.scale = 1.5;
        }
        console.log(diff - this.prevDiff, this.scale);
        this.prevDiff = diff;
        this.draw({ x: 0, y: 0 });
      } else if (touch.touches.length == 1 && this.moved) {
        console.log(touch.touches[0].clientX - this.prevPos.x, touch.touches[0].clientY - this.prevPos.y);
        this.draw({ x: touch.touches[0].clientX - this.prevPos.x, y: touch.touches[0].clientY - this.prevPos.y })
        this.prevPos.x = touch.touches[0].clientX;
        this.prevPos.y = touch.touches[0].clientY;
        this.pressed = false;

      }
    });
    this.canvas.addEventListener("touchend", (touch) => {
      if (touch.changedTouches.length == 2 && this.zoomMobile) {
        this.zoomMobile = false;
        this.prevDiff = 0;
        this.scale = 1;
      } else if (touch.changedTouches.length == 1 && this.moved && !this.pressed) {
        this.moved = false;
        this.prevDiff = 0;
      } else if (touch.changedTouches.length == 1 && this.pressed){
        this.clickHandler({x: touch.changedTouches[0].clientX, y:touch.changedTouches[0].clientY});
        this.prevPos = {x:0, y:0};
        this.moved = false;
        this.pressed = false;
      }
    })
  }

  private clickHandler(coords: {x: number, y:number}) {
    const canvasBoundingBox = this.canvas.getBoundingClientRect();
    const scale = this.ctx.getTransform().a;
    const xTransform = this.ctx.getTransform().e;
    const yTransform = this.ctx.getTransform().f;

    const pixelX = Math.floor((coords.x - canvasBoundingBox.x - xTransform) / scale);
    const pixelY = Math.floor((coords.y - canvasBoundingBox.y - yTransform) / scale);
    window.dispatchEvent(new CustomEvent("pixelClicked", { detail: { x: pixelX, y: pixelY } }));
    console.log(pixelX, pixelY);
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
    if (this.scale != 1) this.scale = 1;
    this.draw({ x: 0, y: 0 });
  }
}

export type Color = [number, number, number, number];

export type CanvasPixels = {
  colors: Color[]
  width: number
  height: number
}
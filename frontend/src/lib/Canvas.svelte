<script lang="ts">
  import { Canvas } from '../assets/pixel-wars/canvas';
  import logo from '../assets/logo.png';
  let canvasElement: HTMLCanvasElement;
  let canvas: Canvas;
  const image = new Image()

  window.addEventListener('load', async () => {
    canvas = new Canvas(canvasElement);
    
    image.src = logo;
    image.onload = () =>
        canvas.ctx.drawImage(image, window.innerWidth / 2 - 252/2, window.innerHeight / 2 - 252/2, 252, 252);
    
  });
</script>

<div class="canvas-container">
  <canvas bind:this={canvasElement} />
</div>

<style>
  .canvas-container {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>

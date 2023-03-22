<script lang="ts">
  import logo from '../assets/logo.png';
  import { loadImage } from '../assets/pixel-wars/utils/loadImage';
  import { initialLoad } from '../assets/pixel-wars/canvas';
  import CanvasElementController from '../assets/pixel-wars/CanvasController';
  let canvasElement: HTMLCanvasElement;
  let canvasController: CanvasElementController;

  window.addEventListener('load', async () => {
    canvasController = new CanvasElementController(canvasElement);
    const image = await loadImage(logo);

    canvasController.ctx.drawImage(
      image,
      window.innerWidth / 2 - 252 / 2,
      window.innerHeight / 2 - 252 / 2,
      252,
      252
    );

    initialLoad(canvasController);
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

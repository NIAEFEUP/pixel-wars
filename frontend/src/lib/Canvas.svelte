<script lang="ts">
  import logo from '../assets/logo.png';
  import { loadImage } from '../assets/pixel-wars/utils/loadImage';
  import { initialLoad } from '../assets/pixel-wars/canvas';
  import CanvasElementController from '../assets/pixel-wars/CanvasController';
  let canvasElement: HTMLCanvasElement;
  let canvasController: CanvasElementController;

  window.addEventListener('load', async () => {
    canvasController = new CanvasElementController(canvasElement);
    initialLoad(canvasController);
  });
</script>

  <canvas bind:this={canvasElement} />

<style>
  canvas {
    width: 100%;
    height: 100%;
  }
</style>

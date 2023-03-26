<script lang="ts">
  import logo from '../assets/logo.png';
  import { loadImage } from '../assets/pixel-wars/utils/loadImage';
  import { initialLoad } from '../assets/pixel-wars/canvas';
  import CanvasElementController from '../assets/pixel-wars/CanvasController';
    import SubscriptionController from '../assets/pixel-wars/SubscriptionController';
  let canvasElement: HTMLCanvasElement;
  let canvasController: CanvasElementController;
  let subscriptionController: SubscriptionController;

  window.addEventListener('load', async () => {
    canvasController = new CanvasElementController(canvasElement);
    await initialLoad(canvasController);
    subscriptionController = new SubscriptionController(canvasController);
  });
</script>
<div id="canvas-container">
  <canvas bind:this={canvasElement} />
</div>
<style>
</style>

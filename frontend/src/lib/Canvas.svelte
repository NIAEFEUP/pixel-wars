<script lang="ts">
  import logo from '../assets/logo.png';
  import { loadImage } from '../assets/pixel-wars/utils/loadImage';
  import { initialLoad } from '../assets/pixel-wars/canvas';
  import CanvasElementController from '../assets/pixel-wars/CanvasController';
  import SubscriptionController from '../assets/pixel-wars/SubscriptionController';
    import CanvasOverlay from './CanvasOverlay.svelte';

  let canvasElement: HTMLCanvasElement;
  let canvasController: CanvasElementController;
  let subscriptionController: SubscriptionController;

  window.addEventListener('load', async () => {
    canvasController = new CanvasElementController(canvasElement);
    await initialLoad(canvasController);
    subscriptionController = new SubscriptionController(canvasController);
    await subscriptionController.initConnection();
  });
</script>
<div id="canvas-container">
  <canvas id="canvas-square" bind:this={canvasElement} />
  <CanvasOverlay />
</div>
<style>
  #canvas-container{
    display: grid;
    grid-template-columns: repeat(1,1fr);
    grid-template-rows: repeat(1,1fr);
  }

  #canvas-square{
    grid-column: 1;
    grid-row: 1;
  }

  
</style>

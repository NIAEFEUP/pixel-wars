<script lang="ts">
  import { TimeoutStore } from '../assets/pixel-wars/stores';

  let diffTime: number = 0;

  TimeoutStore.subscribe(async (timeout) => {
    console.log(timeout);
    if (timeout.remainingPixels == 0) {
      let query = await fetchLastTimeout();
      let nextDate = new Date((query.timeout + 60) * 1000);
      if (query.remainingPixels != 0) {
        TimeoutStore.set({
          timeout: new Date(query.timeout * 1000),
          remainingPixels: query.remainingPixels
        });
        return;
      }
      let timeoutHandle = setInterval(()=>{
        diffTime = (nextDate.getTime() - new Date().getTime())
      }, 1000)
      await new Promise((r) => setTimeout(r, nextDate.getTime() - new Date().getTime()));
      query = await fetchLastTimeout();
      while (query.remainingPixels == 0) {
        await new Promise((r) => setTimeout(r, 1000));
        query = await fetchLastTimeout();
      }
      clearTimeout(timeoutHandle);
      TimeoutStore.set({
        timeout: new Date(query.timeout * 1000),
        remainingPixels: query.remainingPixels
      });
    }
  });

  $: secondsLeft = Math.round(diffTime / 1000);

  async function fetchLastTimeout(): Promise<{ timeout: number; remainingPixels: number }> {
    const query = await fetch('./api/client/details');
    if (query.status != 200) {
      return {timeout: -1, remainingPixels: -1};
    }
    const json = await query.json();
    return { timeout: json.lastTimestamp, remainingPixels: json.remainingPixels };
  }

  window.addEventListener('sessionLoaded', async (ev) => {
    const query = await fetchLastTimeout();
    if(query.remainingPixels == -1){
      TimeoutStore.set({
        timeout: new Date(0),
        remainingPixels: 9
      });
      return;
    }
    TimeoutStore.set({
      timeout: new Date(query.timeout * 1000),
      remainingPixels: query.remainingPixels
    });
  });
</script>

<div id="canvas-overlay" class:canvas-overlay-active={$TimeoutStore.remainingPixels == 0}>
  <div id="canvas-bg" />
  <h1>You changed too many pixels!</h1>
  <h2>Seconds left: {secondsLeft}</h2>
</div>

<style>
  #canvas-overlay {
    position: relative;
    width: 100vw;
    height: 100vh;
    grid-column: 1;
    grid-row: 1;
    z-index: 99;
    display: none;
    background: rgba(150, 150, 150, 0.5);
  }

  #canvas-overlay > h1 {
    text-align: center;
    color: white;
  }
  #canvas-overlay > h2 {
    text-align: center;
    color: white;
  }

  .canvas-overlay-active {
    display: block !important;
  }
</style>

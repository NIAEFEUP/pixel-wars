<script lang="ts">
  import ColorPicker from "./ColorPicker.svelte";
  import Register from './Register.svelte';
  let title: HTMLHeadingElement;
  let formVisible: boolean = false;
  const randomColor = '#' + (((1 << 24) * Math.random()) | 0).toString(16);

  window.addEventListener('load', () => {
    title.style.setProperty('--borderColor', randomColor);
    document.getElementById("avatarButton").addEventListener('click', () => {
      formVisible = !formVisible;
    });
  });
</script>

<div class="navbar">
  <nav>
    <h1 id="title" bind:this={title}>PixelWars</h1>
    <ColorPicker></ColorPicker>
    
    <button id="avatarButton"
      ><img
        class="avatar"
        src="https://source.boringavatars.com/beam/120/{crypto.randomUUID()}colors=001219,05f73,0a9396,94d2bd,e9d8a6,ee9b00,ca6702,bb3e03,ae2012,9b2226"
        alt="your avatar"
      /></button
    >
    {#if formVisible}
        <Register/>
    {/if}
  </nav>
</div>

<style>
  @import url('https://fonts.googleapis.com/css2?family=Bitter:wght@300&family=Kaushan+Script&family=Lobster&family=Major+Mono+Display&family=Ranchers&display=swap');

  .navbar {
    user-select: none;
    width: 100%;
  }

  .navbar nav {
    display: flex;
    align-items: center;
    justify-content: space-between;

    color: var(--vanilla);
  }

  nav > *:nth-child(odd) {
    width: 60px;
  }

  .avatar {
    max-height: 6vh;
  }

  #title {
    line-height: 6vh;
    font-size: 6vh;
    padding: 3vh;

    --borderColor: #ffff;
    -webkit-text-stroke: var(--borderColor) 2px;

    animation-name: fontChange;
    animation-duration: 5s;
    animation-iteration-count: infinite;
  }

  #avatarButton {
    background-color: unset;
    border-style: none;
    cursor: pointer;
    line-height: 20px;
    list-style: none;
    margin: 0;
    outline: none;
    font-size: 2.2vh;
    font-weight: bolder;
    padding: 10px 16px;
    position: relative;
    text-align: center;
    text-decoration: none;
    transition: color 100ms;
    vertical-align: baseline;
    touch-action: manipulation;
  }

  @keyframes fontChange {
    0% {
      font-family: 'Major Mono Display', monospace;
    }
    20% {
      font-family: 'Kaushan Script', cursive;
    }
    40% {
      font-family: 'Lobster', cursive;
    }
    60% {
      font-family: 'Ranchers', cursive;
    }
    80% {
      font-family: 'Bitter', serif;
    }
    100% {
      font-family: 'Major Mono Display', monospace;
    }
  }

  @media (max-width: 768px) {
    #title {
      display: none;
      font-size: 3vh;
      line-height: 4vh;

      -webkit-text-stroke: var(--borderColor) 1px;
    }

    .avatar {
      max-height: 4vh;
    }

    .navbar {
      width: 100%;
    }

    .navbar nav{
      justify-content: space-between;
      padding: 1vh;
    }
  }

  @media (prefers-color-scheme: light) {
    .navbar nav {
      color: var(--dark-cyan);
    }
  }
</style>

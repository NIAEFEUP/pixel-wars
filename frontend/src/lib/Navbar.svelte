<script lang="ts">
  let title: HTMLHeadingElement;
  const randomColor = '#' + (((1 << 24) * Math.random()) | 0).toString(16);

  onload = () => title.style.setProperty('--borderColor', randomColor);
</script>

<div class="navbar">
  <nav>
    <div />

    <h1 id="title" bind:this={title}>PixelWars</h1>

    <a href="/me"
      ><img
        class="avatar"
        src="https://source.boringavatars.com/beam/120/{crypto.randomUUID()}colors=001219,05f73,0a9396,94d2bd,e9d8a6,ee9b00,ca6702,bb3e03,ae2012,9b2226"
        alt="your avatar"
      /></a
    >
  </nav>
</div>

<style>
  @import url('https://fonts.googleapis.com/css2?family=Bitter:wght@300&family=Kaushan+Script&family=Lobster&family=Major+Mono+Display&family=Ranchers&display=swap');

  nav > *:nth-child(odd) {
    width: 60px;
  }

  .navbar {
    width: calc(100% - 10vh);
    padding: 2vh 5vh;
  }

  .navbar nav {
    display: flex;
    align-items: center;
    justify-content: space-between;

    color: var(--vanilla);
  }

  .avatar {
    max-height: 6vh;
  }

  #title {
    line-height: 6vh;
    font-size: 6vh;

    --borderColor: #ffff;
    -webkit-text-stroke: var(--borderColor) 2px;

    animation-name: fontChange;
    animation-duration: 5s;
    animation-iteration-count: infinite;
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
      font-size: 3vh;
      line-height: 4vh;

      -webkit-text-stroke: var(--borderColor) 1px;
    }

    .avatar {
      max-height: 4vh;
    }

    .navbar {
      width: calc(100% - 6vh);
      padding: 2vh 3vh;
    }
  }

  @media (prefers-color-scheme: light) {
    .navbar nav {
    color: var(--dark-cyan);
  }
  }
</style>

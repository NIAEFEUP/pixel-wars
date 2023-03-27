<script lang="ts">
import { ColorPallete } from "../assets/pixel-wars/canvas";
import { ColorPickerStore } from "../assets/pixel-wars/ColorPickerStore";

let number = 0;

const changeColor = num => (mouseEvent:MouseEvent) => {
    number = num;
    const element = mouseEvent.target as Element;
    element.classList.add("color-block-active");
    ColorPickerStore.set(num);
}; 

</script>
<div id="color-picker">
    {#each ColorPallete as color, i}
        <button class="color-block" style="background-color: rgb({color[0]},{color[1]},{color[2]});"
            on:click="{changeColor(i)}"
            class:color-block-active="{i === number}">
        </button>
        
    {/each}
</div>
<style>
    #color-picker{
        display: grid;
        grid-template-columns: repeat(8, 1fr);
    }
    
    .color-block{
        width: 3vw;
        height: 3vw;
        margin: 0.3vh;
        border: 0;
    }

    .color-block-active{
        border: 0.1vh solid red;
    }


    @media (max-width: 768px) {

        .color-block{
            margin: 0.3vh;
            width: 3vh;
            height: 3vh;
        }
    }
</style>
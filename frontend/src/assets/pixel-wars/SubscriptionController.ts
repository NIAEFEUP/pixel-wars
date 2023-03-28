import type CanvasElementController from './CanvasController';
import { decodeColor } from './canvas';
import { ColorPickerStore, TimeoutStore } from './stores';
import { get } from 'svelte/store';

export default class SubscriptionController {
  websocketServer: WebSocket;
  canvasController: CanvasElementController;

  constructor(canvasController: CanvasElementController) {
    this.canvasController = canvasController;
    ColorPickerStore.subscribe((val) => {
      console.log(val);
    });
  }


  public async initConnection() {
    const cookies = await fetch('pixelwars/api/getSession');
    if(cookies.status == 401){
      console.log("Client already has session...");
    }
    this.websocketServer = new WebSocket('wss://'+window.location.host+'/pixelwars/api/subscribe');
    this.websocketServer.addEventListener("message", this.receiveMessageHandler());

    window.addEventListener("pixelClicked", async (ev:CustomEvent) => {
      const coords = ev.detail as {x:number, y:number};
      let timeout = get(TimeoutStore);
      if(timeout.remainingPixels == 0) return 0;
      timeout.remainingPixels--;
      TimeoutStore.set(timeout);
      const color = get(ColorPickerStore);
      await this.sendUpdate(coords.x, coords.y, color);
      this.canvasController.putPixelCanvas(coords.x,coords.y, decodeColor(color));
    })
  }

  public async sendUpdate(x: number, y: number, color: number) {
    if (color >= 16) throw new Error(`illegal color ${color} must be less than 16...`);
    this.websocketServer.send(this.encodeMessage(x, y, color));
  }

  private encodeMessage(x: number, y: number, color: number) {
    const buffer = new ArrayBuffer(5);
    const dataView = new DataView(buffer);

    dataView.setUint16(0, x, false);
    dataView.setUint16(2, y, false);
    dataView.setUint8(4, color);

    return buffer;
  }

  private decodeMessage(buffer: ArrayBuffer) {
    const dataView = new DataView(buffer);

    const x = dataView.getUint16(0, false);
    const y = dataView.getUint16(2, false);
    const color = dataView.getUint8(4);

    return { x, y, color };
  }
  
  private receiveMessageHandler() {
    const subscription:SubscriptionController = this;
    return  async (message: MessageEvent<Blob>) => {
      const { x, y, color } = subscription.decodeMessage(await message.data.arrayBuffer());

      subscription.canvasController.putPixelCanvas(x, y, decodeColor(color));
    };
  }

}

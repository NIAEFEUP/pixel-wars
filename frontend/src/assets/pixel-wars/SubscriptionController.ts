import type CanvasElementController from './CanvasController';
import { decodeColor } from './canvas';

export default class SubscriptionController {
  websocketServer: WebSocket;
  canvasController: CanvasElementController;

  constructor(canvasController: CanvasElementController) {
    this.canvasController = canvasController;
  }

  private async receiveMessageHandler(message: MessageEvent<ArrayBuffer>) {
    const { x, y, color } = this.decodeMessage(message.data);

    this.canvasController.putPixelCanvas(x, y, decodeColor(color));
  }

  public async initConnection(canvasController: CanvasElementController) {
    const cookies = await fetch(window.location.host + '/api/getSession');

    this.websocketServer = new WebSocket('ws://' + window.location.hostname + '/api/subscribe');
    this.websocketServer.onmessage = this.receiveMessageHandler;
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
}

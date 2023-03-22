import type CanvasElementController from "./CanvasController" 

export default class SubscriptionController{

    websocketServer: WebSocket
    canvasController: CanvasElementController

    constructor(canvasController: CanvasElementController){
        this.canvasController = canvasController;
    }

    private async receiveMessageHandler(message: MessageEvent<ArrayBuffer>){
    }

    public async initConnection(){
        const cookies = await fetch(window.location.host + "/api/getSession");

        this.websocketServer = new WebSocket("ws://"+window.location.hostname + "/api/subscribe");  
        this.websocketServer.onmessage = this.receiveMessageHandler;
    }

    public async sendUpdate(x:number, y:number, color: number){

        if(color >= 16) throw new Error(`illegal color ${color} must be less than 16...`);


        const buffer = new ArrayBuffer(5);
        const dataView = new DataView(buffer);

        dataView.setUint16(0, x, false);
        dataView.setUint16(3, y, false);
        dataView.setUint8(4, color);

        this.websocketServer.send(buffer);

    }
}
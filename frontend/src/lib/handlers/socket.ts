import type {Message} from "$lib/models/chat/message/Message";
import {PUBLIC_WS_API} from "$env/static/public";
import {size} from "$lib/stores/size";
import {get} from "svelte/store";

export default class TextSocket {
    private socket: WebSocket;
    private readonly onReceive: (n: number, a: Array<Message>) => void;
    constructor(onReceive: (n: number, a: Array<Message>) => void) {
        this.onReceive = onReceive;
        this.socket = new WebSocket(PUBLIC_WS_API);
        this.socket.onmessage = async(e: MessageEvent) => {
            console.log("received request");
            if (e.data != "null") {
                const elements: Array<Message> = await JSON.parse(e.data);
                await this.onReceive(elements.length, elements);
            }
            this.socket.send(get(size) as string);
        };
        this.socket.onopen = () => {
            console.log("Socket Connected");
            console.log(`Client: ${get(size)}`);
            this.socket.send(get(size) as string);
        };
        this.socket.onclose = () => {
            console.log("Socket Closed");
        };
    }
}
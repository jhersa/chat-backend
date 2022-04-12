import { createApp } from "Vue";

createApp({
  el: "#app",
  data: {
    ws: null,
    serverUrl: "ws://" + window.location.host + "/ws",
  },
  mounted: () => {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
      this.ws = new WebSocket(this.serverUrl);
      this.ws.addEventListener("open", (event) => {
        this.onWebsocketOpen(event);
      });
    },
    onWebsocketOpen() {
      console.log("Conectado al websocket");
    },
  },
});

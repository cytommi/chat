const WS_URL = "ws://localhost:9090/ws";

export const skt = new WebSocket(WS_URL);

skt.onopen = (_) => {
  console.log("Connected to websocket");
};

skt.onclose = (_) => {
  console.log("Disconnected from websocket");
};

skt.onmessage = (ev) => {
  console.log("Message received: ", ev.data);

  const msgs = document.getElementById("message-list");
  // if (msgs.children.length > 10) {
  // msgs.removeChild(msgs.children[0]);
  // }
  msgs.appendChild(makeMessageListItem(ev.data));
};

function makeMessageListItem(msg) {
  const li = document.createElement("li");
  li.textContent = msg;
  return li;
}

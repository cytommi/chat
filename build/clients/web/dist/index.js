const url = "ws://localhost:9090/ws";

const skt = new WebSocket(url);
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

function sendMessage(ev) {
  ev.preventDefault();
  // get values
  const usernameInput = document.getElementById("username-input");
  const msgInput = document.getElementById("message-input");

  const username = usernameInput.value;
  const msg = msgInput.value;

  skt.send(JSON.stringify({ username, msg }));
}

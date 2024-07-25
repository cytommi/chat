import { skt } from "./ws.js";

const MESSAGE_PAYLOAD_TYPE = "MSG";
const AUTH_PAYLOAD_TYPE = "AUTH";

export function sendMessage(ev) {
  console.log("sending");
  ev.preventDefault();
  // get values
  // const usernameInput = document.getElementById("username-input");
  const msgInput = document.getElementById("message-input");

  // const username = usernameInput.value;
  const msg = msgInput.value;

  skt.send(JSON.stringify({ type: MESSAGE_PAYLOAD_TYPE, msg }));
}

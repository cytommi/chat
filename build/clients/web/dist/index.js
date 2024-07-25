import { authenticate } from "./auth.js";
import { sendMessage } from "./message.js";

window.onload = async () => {
  const token = await authenticate("tommy");
  console.log({ token });

  await joinRoom("room_1", token);
  console.log("joined!");
  document.getElementById("send-message").disabled = false;
};

async function joinRoom(roomId, token) {
  const res = await fetch(`http://localhost:9090/room/${roomId}`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${token}`,
    }),
  });
  if (!res.ok) {
    alert("failed to join room");
  }
}

// Expose global functions
globalThis.sendMessage = sendMessage;

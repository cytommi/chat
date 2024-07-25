const AUTHENTICATE_ENDPOINT = "http://localhost:9090/authenticate";

export async function authenticate(userId) {
  const res = await fetch(AUTHENTICATE_ENDPOINT, {
    method: "POST",
    body: JSON.stringify({
      userId,
    }),
  });
  if (!res.ok) {
    alert("failed to authenticate");
  }
  const { token } = await res.json();
  return token;
}

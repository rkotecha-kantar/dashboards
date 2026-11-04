async function request(url, options) {
  const res = await fetch(url, options);
  if (!res.ok) throw new Error(`Request failed (${res.status})`);
  return res.status === 204 ? null : res.json();
}

export const getWidgets = () => request("/api/widgets");

export const createWidget = (widget) =>
  request("/api/widgets", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(widget),
  });

export const deleteWidget = (id) =>
  request(`/api/widgets/${id}`, { method: "DELETE" });

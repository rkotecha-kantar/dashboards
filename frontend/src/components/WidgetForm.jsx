import { useState } from "react";

export default function WidgetForm({ onAdd }) {
  const [title, setTitle] = useState("");
  const [value, setValue] = useState("");

  const submit = (e) => {
    e.preventDefault();
    if (!title.trim()) return;
    onAdd({ title, value: Number(value) || 0 });
    setTitle("");
    setValue("");
  };

  return (
    <form onSubmit={submit} style={{ margin: "16px 0" }}>
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Title"
      />{" "}
      <input
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Value"
        type="number"
      />{" "}
      <button type="submit">Add</button>
    </form>
  );
}

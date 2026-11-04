export default function WidgetCard({ widget, onDelete }) {
  return (
    <div
      style={{
        border: "1px solid #ccc",
        borderRadius: 8,
        padding: 16,
        flex: 1,
      }}
    >
      <div style={{ color: "#666" }}>{widget.title}</div>
      <div style={{ fontSize: 32, fontWeight: "bold" }}>{widget.value}</div>
      <button onClick={() => onDelete(widget.id)}>Delete</button>
    </div>
  );
}

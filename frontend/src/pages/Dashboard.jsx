import WidgetCard from "../components/WidgetCard";
import WidgetForm from "../components/WidgetForm";
import { useWidgets } from "../hooks/useWidgets";

export default function Dashboard({
  widgets: initialWidgets,
  userName,
  appName,
}) {
  const { widgets, error, add, remove } = useWidgets(initialWidgets);

  return (
    <div
      style={{ maxWidth: 720, margin: "40px auto", fontFamily: "sans-serif" }}
    >
      <h1>{appName}</h1>
      <p>Welcome back, {userName}.</p>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <WidgetForm onAdd={add} />
      <div style={{ display: "flex", gap: 16, flexWrap: "wrap" }}>
        {widgets.map((w) => (
          <WidgetCard key={w.id} widget={w} onDelete={remove} />
        ))}
      </div>
    </div>
  );
}

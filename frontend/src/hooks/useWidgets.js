import { useState } from "react";
import { getWidgets, createWidget, deleteWidget } from "../api/widgets";

export function useWidgets(initialWidgets = []) {
  const [widgets, setWidgets] = useState(initialWidgets);
  const [error, setError] = useState(null);

  const refresh = async () => {
    try {
      setWidgets(await getWidgets());
      setError(null);
    } catch (e) {
      setError(e.message);
    }
  };

  const add = async (widget) => {
    try {
      await createWidget(widget);
      await refresh();
    } catch (e) {
      setError(e.message);
    }
  };

  const remove = async (id) => {
    try {
      await deleteWidget(id);
      await refresh();
    } catch (e) {
      setError(e.message);
    }
  };

  return { widgets, error, add, remove };
}

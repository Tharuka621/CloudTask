import { useState } from 'react';
import { Plus } from 'lucide-react';

export default function Dashboard() {
  const [tasks] = useState([
    { id: 1, title: 'Setup microservices', status: 'TODO', priority: 'HIGH' },
    { id: 2, title: 'Implement Auth', status: 'IN_PROGRESS', priority: 'HIGH' },
    { id: 3, title: 'Create DB schema', status: 'DONE', priority: 'MEDIUM' },
  ]);

  const columns = [
    { id: 'TODO', title: 'To Do' },
    { id: 'IN_PROGRESS', title: 'In Progress' },
    { id: 'DONE', title: 'Completed' },
  ];

  const getPriorityColor = (p: string) => {
    switch (p) {
      case 'HIGH': return '#ef4444';
      case 'MEDIUM': return '#f59e0b';
      case 'LOW': return '#10b981';
      default: return '#6b7280';
    }
  };

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '32px' }}>
        <div>
          <h1 className="text-h1">Engineering Team</h1>
          <p className="text-subtle" style={{ marginTop: '8px' }}>Manage and track your software projects here.</p>
        </div>
        <button className="btn btn-primary">
          <Plus size={18} /> New Task
        </button>
      </div>

      <div className="kanban-board">
        {columns.map(col => (
          <div key={col.id} className="kanban-column">
            <div className="kanban-column-header">
              <h3 className="text-h3" style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                {col.title}
                <span style={{ fontSize: '12px', background: 'var(--border-color)', padding: '2px 8px', borderRadius: '12px' }}>
                  {tasks.filter(t => t.status === col.id).length}
                </span>
              </h3>
            </div>
            <div className="kanban-cards">
              {tasks.filter(t => t.status === col.id).map(task => (
                <div key={task.id} className="task-card" draggable>
                  <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '8px' }}>
                    <span style={{ fontSize: '10px', fontWeight: 700, letterSpacing: '0.05em', color: getPriorityColor(task.priority), padding: '2px 6px', background: `${getPriorityColor(task.priority)}15`, borderRadius: '4px' }}>
                      {task.priority}
                    </span>
                    <span className="text-xs text-subtle">TASK-{task.id}</span>
                  </div>
                  <h4 style={{ fontSize: '14px', fontWeight: 500, lineHeight: 1.4 }}>{task.title}</h4>
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

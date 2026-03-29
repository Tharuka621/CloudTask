import { NavLink } from 'react-router-dom';
import { Home, CheckSquare, Users, Settings } from 'lucide-react';

export default function Sidebar() {
  const navItems = [
    { name: 'Dashboard', icon: Home, path: '/' },
    { name: 'My Tasks', icon: CheckSquare, path: '/tasks' },
    { name: 'Teams', icon: Users, path: '/teams' },
    { name: 'Settings', icon: Settings, path: '/settings' },
  ];

  return (
    <aside className="sidebar">
      <div style={{ padding: '24px', display: 'flex', alignItems: 'center', gap: '12px' }}>
        <div style={{ width: '32px', height: '32px', backgroundColor: 'var(--accent-primary)', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'white', fontWeight: 'bold' }}>
          CT
        </div>
        <span className="text-h2">CloudTask</span>
      </div>
      <nav style={{ flex: 1, paddingTop: '16px' }}>
        {navItems.map((item) => (
          <NavLink
            key={item.name}
            to={item.path}
            className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}
          >
            <item.icon size={18} />
            <span>{item.name}</span>
          </NavLink>
        ))}
      </nav>
      <div style={{ padding: '24px', borderTop: '1px solid var(--border-color)' }}>
        <p className="text-xs text-subtle">CloudTask v1.0.0</p>
      </div>
    </aside>
  );
}

import { Bell, Search, User } from 'lucide-react';
import { useAuthStore } from '../store/authStore';

export default function Header() {
  const { user, logout } = useAuthStore();

  return (
    <header className="header" style={{ justifyContent: 'space-between' }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: '8px', width: '300px' }}>
        <div className="input" style={{ display: 'flex', alignItems: 'center', gap: '8px', padding: '6px 12px' }}>
          <Search size={16} className="text-subtle" />
          <input type="text" placeholder="Search..." style={{ border: 'none', outline: 'none', background: 'transparent', width: '100%' }} />
        </div>
      </div>
      <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
        <button className="btn-ghost" style={{ borderRadius: '50%', padding: '8px' }}>
          <Bell size={20} />
        </button>
        <div style={{ display: 'flex', alignItems: 'center', gap: '12px', borderLeft: '1px solid var(--border-color)', paddingLeft: '16px' }}>
          <div style={{ textAlign: 'right' }}>
            <p className="text-sm" style={{ fontWeight: 600 }}>{user?.email || 'User'}</p>
            <p className="text-xs text-subtle">{user?.role || 'Member'}</p>
          </div>
          <button onClick={logout} className="btn-ghost" style={{ borderRadius: '50%', padding: '8px', backgroundColor: 'var(--bg-hover)' }}>
            <User size={20} />
          </button>
        </div>
      </div>
    </header>
  );
}

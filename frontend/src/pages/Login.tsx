import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useAuthStore } from '../store/authStore';
import api from '../lib/api';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const setAuth = useAuthStore((state) => state.setAuth);
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/auth/login', { email, password });
      const { access_token, refresh_token } = res.data;
      // Decode the JWT to get the actual role assigned by the database
      const payload = JSON.parse(atob(access_token.split('.')[1]));
      const user = { id: payload.user_id, email, role: payload.role };
      setAuth(user, access_token, refresh_token);
      navigate('/');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Login failed');
    }
  };

  return (
    <div style={{ display: 'flex', height: '100vh', alignItems: 'center', justifyContent: 'center', backgroundColor: 'var(--bg-secondary)' }}>
      <div className="card" style={{ width: '400px' }}>
        <div style={{ textAlign: 'center', marginBottom: '24px' }}>
          <div style={{ width: '48px', height: '48px', backgroundColor: 'var(--accent-primary)', borderRadius: '12px', display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'white', fontWeight: 'bold', margin: '0 auto 16px', fontSize: '20px' }}>
            CT
          </div>
          <h1 className="text-h1">Welcome back</h1>
          <p className="text-subtle" style={{ marginTop: '8px' }}>Enter your details to access your tasks.</p>
        </div>
        
        <form onSubmit={handleLogin} style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
          {error && <div style={{ color: 'red', fontSize: '13px', backgroundColor: '#fee2e2', padding: '10px', borderRadius: '6px' }}>{error}</div>}
          
          <div>
            <label className="label">Email</label>
            <input type="email" className="input" value={email} onChange={(e) => setEmail(e.target.value)} required placeholder="name@company.com" />
          </div>
          <div>
            <label className="label">Password</label>
            <input type="password" className="input" value={password} onChange={(e) => setPassword(e.target.value)} required placeholder="••••••••" />
          </div>
          
          <button type="submit" className="btn btn-primary" style={{ marginTop: '8px', padding: '10px' }}>
            Sign In
          </button>
        </form>

        <p style={{ textAlign: 'center', marginTop: '24px', fontSize: '14px', color: 'var(--text-secondary)' }}>
          Don't have an account? <Link to="/register" className="text-accent" style={{ fontWeight: 500 }}>Sign up</Link>
        </p>
      </div>
    </div>
  );
}

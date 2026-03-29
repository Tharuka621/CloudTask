import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import api from '../lib/api';

export default function Register() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.post('/auth/register', { email, password });
      navigate('/login');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Registration failed');
    }
  };

  return (
    <div style={{ display: 'flex', height: '100vh', alignItems: 'center', justifyContent: 'center', backgroundColor: 'var(--bg-secondary)' }}>
      <div className="card" style={{ width: '400px' }}>
        <div style={{ textAlign: 'center', marginBottom: '24px' }}>
          <h1 className="text-h1">Create an account</h1>
          <p className="text-subtle" style={{ marginTop: '8px' }}>Start managing your tasks today.</p>
        </div>
        
        <form onSubmit={handleRegister} style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
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
            Get Started
          </button>
        </form>

        <p style={{ textAlign: 'center', marginTop: '24px', fontSize: '14px', color: 'var(--text-secondary)' }}>
          Already have an account? <Link to="/login" className="text-accent" style={{ fontWeight: 500 }}>Sign in</Link>
        </p>
      </div>
    </div>
  );
}

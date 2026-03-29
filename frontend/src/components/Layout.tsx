import { Outlet, Navigate } from 'react-router-dom';
import Sidebar from './Sidebar';
import Header from './Header';
import { useAuthStore } from '../store/authStore';

export default function Layout() {
  const accessToken = useAuthStore((state) => state.accessToken);

  if (!accessToken) {
    return <Navigate to="/login" replace />;
  }

  return (
    <div className="app-layout">
      <Sidebar />
      <main className="main-content">
        <Header />
        <div className="page-container">
          <Outlet />
        </div>
      </main>
    </div>
  );
}

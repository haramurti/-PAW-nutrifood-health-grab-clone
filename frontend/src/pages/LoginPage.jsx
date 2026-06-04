import { useState } from 'react';
import { useDispatch } from 'react-redux';
import '../styles/LoginPage.css';
import { setRole, setToken } from '../redux/slices/loginSlice';
import { useNavigate } from 'react-router-dom';

const AUTH_URL = 'http://localhost:8001/api/auth';

function LoginPage() {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const [mode, setMode] = useState('login'); // 'login' | 'register'
  const [role, setRoleState] = useState('customer'); // 'customer' | 'merchant'
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleLogin = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      const res = await fetch(`${AUTH_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      const data = await res.json();
      if (!res.ok) {
        if (data.errors) {
          const firstError = Object.values(data.errors)[0][0];
          throw new Error(firstError);
        }
        throw new Error(data.message || 'Login gagal.');
      }
      
      dispatch(setToken(data.data.token));
      dispatch(setRole(role));
      navigate('/products');
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    if (password !== passwordConfirm) {
      setError('Konfirmasi password tidak cocok.');
      return;
    }
    setLoading(true);
    setError('');
    try {
      const res = await fetch(`${AUTH_URL}/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name,
          email,
          password,
          password_confirmation: passwordConfirm,
        }),
      });
      const data = await res.json();
      if (!res.ok) {
        if (data.errors) {
          const firstError = Object.values(data.errors)[0][0];
          throw new Error(firstError);
        }
        throw new Error(data.message || 'Register gagal.');
      }
      
      dispatch(setToken(data.data.token));
      dispatch(setRole(role));
      navigate('/products');
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      {/* Logo Section */}
      <div className="login-logo-section">
        <div className="login-logo-icon">
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10c1.85 0 3.58-.5 5.07-1.38l-1.45-1.73C14.44 19.59 13.26 20 12 20c-4.42 0-8-3.58-8-8s3.58-8 8-8c3.18 0 5.93 1.87 7.21 4.56L21.5 7.5C19.79 4.25 16.16 2 12 2z" fill="white"/>
            <path d="M17 8c-1.5 0-3.5 1.5-4 3-.5-1.5-2.5-3-4-3-2.5 0-4 2-4 4.5S8 17 12 20c4-3 7-5 7-7.5S19.5 8 17 8z" fill="white" opacity="0.8"/>
            <path d="M15 3c0 0 1 2 1 4s-1 3-1 3" stroke="white" strokeWidth="1.5" strokeLinecap="round"/>
            <path d="M18 4c0 0 1 1.5 1 3s-1 2.5-1 2.5" stroke="white" strokeWidth="1.5" strokeLinecap="round" opacity="0.7"/>
          </svg>
        </div>
        <h1 className="login-title">Healthy</h1>
        <p className="login-subtitle">Your Nutrition Partner</p>
      </div>

      {/* Login Card */}
      <div className="login-card">
        {/* Role Toggle */}
        <div className="role-toggle">
          <button
            className={`role-toggle-btn ${role === 'customer' ? 'active' : ''}`}
            onClick={() => setRoleState('customer')}
            type="button"
          >
            Customer
          </button>
          <button
            className={`role-toggle-btn ${role === 'merchant' ? 'active' : ''}`}
            onClick={() => setRoleState('merchant')}
            type="button"
          >
            Merchant
          </button>
        </div>

        {/* Error Alert */}
        {error && (
          <div className="bg-red-50 border border-red-200 text-red-600 text-xs px-4 py-2 rounded-lg mb-4">
            {error}
          </div>
        )}

        {/* Form */}
        <form onSubmit={mode === 'login' ? handleLogin : handleRegister}>
          {mode === 'register' && (
            <div className="login-input-group">
              <label>Nama Lengkap</label>
              <div className="login-input-wrapper">
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" fill="currentColor"/>
                </svg>
                <input
                  type="text"
                  placeholder="Masukkan nama lengkap"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  required
                />
              </div>
            </div>
          )}

          <div className="login-input-group">
            <label>Alamat Email</label>
            <div className="login-input-wrapper">
              <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 4H4C2.9 4 2 4.9 2 6V18C2 19.1 2.9 20 4 20H20C21.1 20 22 19.1 22 18V6C22 4.9 21.1 4 20 4ZM20 8L12 13L4 8V6L12 11L20 6V8Z" fill="currentColor"/>
              </svg>
              <input
                type="email"
                placeholder="Masukkan email Anda"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </div>
          </div>

          <div className="login-input-group">
            <label>Password</label>
            <div className="login-input-wrapper">
              <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M18 8H17V6C17 3.24 14.76 1 12 1C9.24 1 7 3.24 7 6V8H6C4.9 8 4 8.9 4 10V20C4 21.1 4.9 22 6 22H18C19.1 22 20 21.1 20 20V10C20 8.9 19.1 8 18 8ZM12 17C10.9 17 10 16.1 10 15C10 13.9 10.9 13 12 13C13.1 13 14 13.9 14 15C14 16.1 13.1 17 12 17ZM15.1 8H8.9V6C8.9 4.29 10.29 2.9 12 2.9C13.71 2.9 15.1 4.29 15.1 6V8Z" fill="currentColor"/>
              </svg>
              <input
                type={showPassword ? 'text' : 'password'}
                placeholder="Masukkan password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
              />
              <button
                type="button"
                className="password-toggle-btn"
                onClick={() => setShowPassword(!showPassword)}
              >
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" width="18" height="18">
                  {showPassword ? (
                    <path d="M12 4.5C7 4.5 2.73 7.61 1 12C2.73 16.39 7 19.5 12 19.5C17 19.5 21.27 16.39 23 12C21.27 7.61 17 4.5 12 4.5ZM12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17ZM12 9C10.34 9 9 10.34 9 12C9 13.66 10.34 15 12 15C13.66 15 15 13.66 15 12C15 10.34 13.66 9 12 9Z" fill="currentColor"/>
                  ) : (
                    <path d="M12 7C14.76 7 17 9.24 17 12C17 12.65 16.87 13.26 16.64 13.83L19.56 16.75C21.07 15.49 22.26 13.86 23 12C21.27 7.61 17 4.5 12 4.5C10.6 4.5 9.26 4.75 8.01 5.2L10.17 7.36C10.74 7.13 11.35 7 12 7ZM2 4.27L4.28 6.55L4.74 7.01C3.08 8.3 1.78 10.02 1 12C2.73 16.39 7 19.5 12 19.5C13.55 19.5 15.03 19.2 16.38 18.66L16.81 19.09L19.73 22L21 20.73L3.27 3L2 4.27ZM7.53 9.8L9.08 11.35C9.03 11.56 9 11.78 9 12C9 13.66 10.34 15 12 15C12.22 15 12.44 14.97 12.65 14.92L14.2 16.47C13.53 16.8 12.79 17 12 17C9.24 17 7 14.76 7 12C7 11.21 7.2 10.47 7.53 9.8ZM11.84 9.02L14.99 12.17L15.01 12.01C15.01 10.35 13.67 9.01 12.01 9.01L11.84 9.02Z" fill="currentColor"/>
                  )}
                </svg>
              </button>
            </div>
          </div>

          {mode === 'register' && (
            <div className="login-input-group">
              <label>Konfirmasi Password</label>
              <div className="login-input-wrapper">
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M18 8H17V6C17 3.24 14.76 1 12 1C9.24 1 7 3.24 7 6V8H6C4.9 8 4 8.9 4 10V20C4 21.1 4.9 22 6 22H18C19.1 22 20 21.1 20 20V10C20 8.9 19.1 8 18 8ZM12 17C10.9 17 10 16.1 10 15C10 13.9 10.9 13 12 13C13.1 13 14 13.9 14 15C14 16.1 13.1 17 12 17ZM15.1 8H8.9V6C8.9 4.29 10.29 2.9 12 2.9C13.71 2.9 15.1 4.29 15.1 6V8Z" fill="currentColor"/>
                </svg>
                <input
                  type="password"
                  placeholder="Konfirmasi password"
                  value={passwordConfirm}
                  onChange={(e) => setPasswordConfirm(e.target.value)}
                  required
                />
              </div>
            </div>
          )}

          <button type="submit" disabled={loading} className="login-submit-btn">
            {loading ? (
              <span className="flex items-center justify-center gap-2">
                <span className="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></span>
                Memproses...
              </span>
            ) : mode === 'login' ? 'Login' : 'Register'}
          </button>
        </form>

        {mode === 'login' ? (
          <p className="login-register-link">
            Belum punya akun?
            <a href="#" onClick={(e) => { e.preventDefault(); setMode('register'); setError(''); }}>
              Daftar Sekarang
            </a>
          </p>
        ) : (
          <p className="login-register-link">
            Sudah punya akun?
            <a href="#" onClick={(e) => { e.preventDefault(); setMode('login'); setError(''); }}>
              Login
            </a>
          </p>
        )}
      </div>

      {/* Footer */}
      <p className="login-footer">by Newbiers</p>
    </div>
  );
}

export default LoginPage;
import { createSlice } from "@reduxjs/toolkit";

const loginSlice = createSlice({
  name: 'login',
  initialState: {
    role: localStorage.getItem('role') || '',
    token: localStorage.getItem('jwt_token') || '',
  },
  reducers: {
    setRole(state, action) {
      state.role = action.payload;
      localStorage.setItem('role', action.payload);
    },
    setToken(state, action) {
      state.token = action.payload;
      localStorage.setItem('jwt_token', action.payload);
    },
    logout(state) {
      state.role = '';
      state.token = '';
      localStorage.removeItem('jwt_token');
      localStorage.removeItem('role');
    },
  }
});

export const { setRole, setToken, logout } = loginSlice.actions;

export const login = (isMerchant) => {
  const role = isMerchant ? 'merchant' : 'customer';
  return setRole(role);
};

export default loginSlice.reducer;
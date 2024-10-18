// authActions.js
import { useAuth } from './AuthContext';

export const useAuthActions = () => {
  const { login, logout } = useAuth();

  const performLogin = (user) => {
    login(user);
  };

  const performLogout = () => {
    logout();
  };

  return { performLogin, performLogout };
};

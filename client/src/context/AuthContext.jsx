import React, { createContext, useContext, useReducer, useEffect } from 'react';
import Cookies from 'js-cookie';

const initialState = {
  isLoggedIn: false,
  user: null,
};

const AuthContext = createContext(initialState);

const LOGIN = 'LOGIN';
const LOGOUT = 'LOGOUT';

// Create a reducer to handle authentication actions
const authReducer = (state, action) => {
  switch (action.type) {
    case LOGIN:
      return { ...state, isLoggedIn: true, user: action.payload };
    case LOGOUT:
      return { ...state, isLoggedIn: false, user: null };
    default:
      return state;
  }
};

// Provider component
export const AuthProvider = ({ children }) => {
  const [state, dispatch] = useReducer(authReducer, initialState);
  
  useEffect(() => {
    const storedUser = Cookies.get('user');
    if (storedUser) {
      dispatch({ type: LOGIN, payload: JSON.parse(storedUser) });
    }
  }, []);

  const login = (user) => {
    Cookies.set('user', JSON.stringify(user), { expires: 1 });
    dispatch({ type: LOGIN, payload: user });
  };
  

  const logout = () => {
    Cookies.remove('user');
    dispatch({ type: LOGOUT });
  };
  

  return (
    <AuthContext.Provider value={{ state, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// Custom hook to use the AuthContext
export const useAuth = () => {
  return useContext(AuthContext);
};

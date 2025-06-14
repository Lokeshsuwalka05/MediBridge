import React, { createContext, useContext, useState, useEffect, type ReactNode } from 'react';
import type { User, ApiResponse } from '../types';
import { authService } from '../services/authService';
import api from '../services/api';

interface AuthContextType {
  user: User | null;
  login: (email: string, password: string) => Promise<ApiResponse<User>>;
  logout: () => void;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const initializeAuth = async () => {
      const token = localStorage.getItem('token');
      const userData = localStorage.getItem('user');

      if (token && userData) {
        try {
          // Set the token in API headers
          api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

          // Validate token by making a test request
          const response = await api.get('/auth/validate');

          // If validation succeeds, set the user from the response
          if (response.data && response.data.user) {
            setUser(response.data.user);
          } else {
            // If no user data in response, use stored data
            const parsedUser = JSON.parse(userData);
            setUser(parsedUser);
          }
        } catch (error) {
          console.error('Token validation failed:', error);
          // Clear invalid token and user data
          localStorage.removeItem('token');
          localStorage.removeItem('user');
          delete api.defaults.headers.common['Authorization'];
          setUser(null);
        }
      } else {
        // Clear any existing auth data
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        delete api.defaults.headers.common['Authorization'];
        setUser(null);
      }
      setIsLoading(false);
    };

    initializeAuth();
  }, []);

  const login = async (email: string, password: string): Promise<ApiResponse<User>> => {
    try {
      const response = await authService.login({ email, password });
      if (response.success && response.data) {
        setUser(response.data);
        // Set the token in API headers
        const token = localStorage.getItem('token');
        if (token) {
          api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        }
      }
      return response;
    } catch (error) {
      console.error('Login error in context:', error);
      return {
        success: false,
        data: null as any,
        error: error instanceof Error ? error.message : 'Login failed'
      };
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    delete api.defaults.headers.common['Authorization'];
    setUser(null);
  };

  const value = {
    user,
    login,
    logout,
    isLoading,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

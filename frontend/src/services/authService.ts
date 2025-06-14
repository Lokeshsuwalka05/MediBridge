import api from './api';
import type { LoginCredentials, User, ApiResponse, LoginResponse } from '../types';

export const authService = {
  login: async (credentials: LoginCredentials): Promise<ApiResponse<User>> => {
    try {
      console.log('Attempting login with credentials:', credentials);

      const response = await api.post<LoginResponse>('/login', credentials);
      console.log('Login response:', response.data);

      if (!response.data || !response.data.token || !response.data.user) {
        console.error('Invalid response structure:', response.data);
        return {
          success: false,
          data: null as any,
          error: 'Invalid response from server'
        };
      }

      // Store token and user data
      localStorage.setItem('token', response.data.token);
      localStorage.setItem('user', JSON.stringify(response.data.user));

      // Set the token in the API instance for future requests
      api.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`;

      return {
        success: true,
        data: response.data.user,
        message: 'Login successful'
      };
    } catch (error: any) {
      console.error('Login error details:', {
        status: error.response?.status,
        data: error.response?.data,
        message: error.message,
        error: error,
        config: error.config
      });

      // Handle specific error cases
      if (error.response?.status === 401) {
        return {
          success: false,
          data: null as any,
          error: error.response?.data?.error || 'Invalid credentials'
        };
      }

      if (error.message === 'Network Error') {
        return {
          success: false,
          data: null as any,
          error: 'Network error. Please check your connection.'
        };
      }

      return {
        success: false,
        data: null as any,
        error: error.response?.data?.error || error.message || 'Login failed'
      };
    }
  }
};
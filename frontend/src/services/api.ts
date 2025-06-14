import axios from 'axios';
import { toast } from 'sonner';

// Create axios instance with base URL from environment variables
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
      console.log('Setting Authorization header:', config.headers.Authorization);
    } else {
      console.warn('No token found in localStorage');
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('Response error:', error.response);

    if (error.response) {
      const { status, data } = error.response;

      // Handle 401 Unauthorized
      if (status === 401) {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        toast.error('Your session has expired. Please log in again.');
        return Promise.reject(error);
      }

      // Handle 403 Forbidden
      if (status === 403) {
        const user = JSON.parse(localStorage.getItem('user') || '{}');
        if (user.role === 'doctor') {
          window.location.href = '/doctor/patients';
        } else if (user.role === 'receptionist') {
          window.location.href = '/receptionist/patients';
        }
        toast.error('You do not have permission to access this resource');
        return Promise.reject(error);
      }

      // Handle other errors
      const errorMessage = data?.message || data?.error || 'An error occurred';
      toast.error(errorMessage);
    } else {
      toast.error('Network error. Please check your connection.');
    }

    return Promise.reject(error);
  }
);

export default api;

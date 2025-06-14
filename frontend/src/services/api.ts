import axios from 'axios';
import { toast } from 'sonner';

// Use environment variable for API URL in production
const API_BASE_URL = 'https://medibridge.onrender.com';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
  // Add timeout to prevent hanging requests
  timeout: 10000,
});

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    console.error('Request error:', error);
    return Promise.reject(error);
  }
);

// Response interceptor for error handling
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error('Response error:', error.response || error);

    // Handle authentication errors
    if (error.response?.status === 401) {
      // Only redirect to login if not already on login page
      if (!window.location.pathname.includes('/login')) {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        window.location.href = '/login';
        toast.error('Session expired. Please login again.');
      } else {
        toast.error(error.response?.data?.error || 'Invalid credentials');
      }
    }
    // Handle not found errors
    else if (error.response?.status === 404) {
      toast.error('Resource not found');
    }
    // Handle server errors
    else if (error.response?.status >= 500) {
      toast.error('Server error. Please try again later.');
    }
    // Handle validation errors
    else if (error.response?.data?.error) {
      toast.error(error.response.data.error);
    }
    // Handle timeout errors
    else if (error.code === 'ECONNABORTED') {
      toast.error('Request timed out. Please try again.');
    }
    // Handle network errors
    else if (error.message === 'Network Error') {
      toast.error('Network error. Please check your connection.');
    }
    // Handle other errors
    else {
      toast.error('An unexpected error occurred. Please try again.');
    }
    return Promise.reject(error);
  }
);

export default api;
import React from 'react';
import { Toaster } from 'sonner';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import ProtectedRoute from './components/auth/ProtectedRoute';
import DashboardLayout from './components/layout/DashboardLayout';
import Login from './pages/Login';
import PatientManagement from './pages/receptionist/PatientManagement';
import DoctorPatients from './pages/doctor/DoctorPatients';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      staleTime: 5 * 60 * 1000, // 5 minutes
    },
  },
});

const App: React.FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <BrowserRouter>
          <div className="App">
            <Toaster
              position="top-right"
              richColors
              closeButton
            />
            
            <Routes>
              <Route path="/login" element={<Login />} />
              
              <Route
                path="/"
                element={
                  <ProtectedRoute>
                    <DashboardLayout />
                  </ProtectedRoute>
                }
              >
                {/* Receptionist Routes */}
                <Route
                  path="receptionist/patients"
                  element={
                    <ProtectedRoute allowedRoles={['receptionist']}>
                      <PatientManagement />
                    </ProtectedRoute>
                  }
                />
                
                {/* Doctor Routes */}
                <Route
                  path="doctor/patients"
                  element={
                    <ProtectedRoute allowedRoles={['doctor']}>
                      <DoctorPatients />
                    </ProtectedRoute>
                  }
                />
                
                {/* Default redirects */}
                <Route
                  path=""
                  element={<Navigate to="/login" replace />}
                />
              </Route>
              
              {/* Catch all route */}
              <Route path="*" element={<Navigate to="/login" replace />} />
            </Routes>
          </div>
        </BrowserRouter>
      </AuthProvider>
    </QueryClientProvider>
  );
};

export default App;
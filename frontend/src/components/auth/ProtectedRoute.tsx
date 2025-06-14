import React, { useEffect } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import LoadingSpinner from '../common/LoadingSpinner';

interface ProtectedRouteProps {
  children: React.ReactNode;
  allowedRoles?: ('doctor' | 'receptionist')[];
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children, allowedRoles }) => {
  const { user, isLoading } = useAuth();
  const location = useLocation();

  // Add a small delay to prevent flash of loading state
  const [showLoading, setShowLoading] = React.useState(true);
  useEffect(() => {
    const timer = setTimeout(() => {
      setShowLoading(false);
    }, 1000);
    return () => clearTimeout(timer);
  }, []);

  if (isLoading || showLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <LoadingSpinner size="lg" />
      </div>
    );
  }

  if (!user) {
    // Store the attempted URL for redirecting back after login
    return <Navigate to="/login" state={{ from: location.pathname }} replace />;
  }

  // If we're at the root path, redirect based on role
  if (location.pathname === '/') {
    const redirectPath = user.role === 'receptionist' ? '/receptionist/patients' : '/doctor/patients';
    return <Navigate to={redirectPath} replace />;
  }

  // Check if the user has the required role for this route
  if (allowedRoles && !allowedRoles.includes(user.role)) {
    // Redirect to the appropriate dashboard based on role
    const redirectPath = user.role === 'receptionist' ? '/receptionist/patients' : '/doctor/patients';
    return <Navigate to={redirectPath} replace />;
  }

  return <>{children}</>;
};

export default ProtectedRoute;
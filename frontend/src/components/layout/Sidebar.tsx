import React from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { 
  Users, 
  UserCheck, 
  FileText, 
  Calendar, 
  Settings, 
  LogOut,
  Stethoscope,
  Heart
} from 'lucide-react';
import { useAuth } from '../../context/AuthContext';

const Sidebar: React.FC = () => {
  const { user, logout } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const receptionistMenuItems = [
    { icon: Users, label: 'Patients', path: '/receptionist/patients' },
    { icon: Calendar, label: 'Appointments', path: '/receptionist/appointments' },
    { icon: FileText, label: 'Reports', path: '/receptionist/reports' },
    { icon: Settings, label: 'Settings', path: '/receptionist/settings' },
  ];

  const doctorMenuItems = [
    { icon: UserCheck, label: 'My Patients', path: '/doctor/patients' },
    { icon: FileText, label: 'Medical Records', path: '/doctor/records' },
    { icon: Calendar, label: 'Schedule', path: '/doctor/schedule' },
    { icon: Settings, label: 'Settings', path: '/doctor/settings' },
  ];

  const menuItems = user?.role === 'receptionist' ? receptionistMenuItems : doctorMenuItems;

  return (
    <div className="bg-white h-full w-64 shadow-lg border-r border-gray-200 flex flex-col">
      {/* Logo and Brand */}
      <div className="p-6 border-b border-gray-200">
        <div className="flex items-center space-x-3">
          <div className="bg-primary-500 p-2 rounded-lg">
            <Heart className="w-6 h-6 text-white" />
          </div>
          <div>
            <h1 className="text-xl font-bold text-gray-900">MediBridge</h1>
            <p className="text-sm text-gray-500">Healthcare Management</p>
          </div>
        </div>
      </div>

      {/* User Info */}
      <div className="p-4 border-b border-gray-200">
        <div className="flex items-center space-x-3">
          <div className="bg-primary-100 p-2 rounded-full">
            <Stethoscope className="w-5 h-5 text-primary-600" />
          </div>
          <div>
            <p className="font-medium text-gray-900">{user?.name}</p>
            <p className="text-sm text-gray-500 capitalize">{user?.role}</p>
          </div>
        </div>
      </div>

      {/* Navigation Menu */}
      {/* <nav className="flex-1 p-4">
        <ul className="space-y-2">
          {menuItems.map((item) => {
            const isActive = location.pathname === item.path;
            return (
              <li key={item.path}>
                <Link
                  to={item.path}
                  className={`flex items-center space-x-3 px-4 py-3 rounded-lg transition-all duration-200 ${
                    isActive
                      ? 'bg-primary-50 text-primary-700 border-l-4 border-primary-500'
                      : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'
                  }`}
                >
                  <item.icon className={`w-5 h-5 ${isActive ? 'text-primary-600' : 'text-gray-400'}`} />
                  <span className="font-medium">{item.label}</span>
                </Link>
              </li>
            );
          })}
        </ul>
      </nav> */}

      {/* Logout Button */}
      <div className="p-4 border-t border-gray-200">
        <button
          onClick={handleLogout}
          className="flex items-center space-x-3 w-full px-4 py-3 text-gray-600 hover:bg-red-50 hover:text-red-600 rounded-lg transition-all duration-200"
        >
          <LogOut className="w-5 h-5" />
          <span className="font-medium">Logout</span>
        </button>
      </div>
    </div>
  );
};

export default Sidebar;
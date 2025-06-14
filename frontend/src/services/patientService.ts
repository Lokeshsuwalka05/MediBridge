import api from './api';
import type { Patient, PaginatedResponse } from '../types';

interface GetPatientsParams {
  page: number;
  limit: number;
  search?: string;
}

interface UpdatePatientData {
  firstName?: string;
  lastName?: string;
  email?: string;
  phone?: string;
  dateOfBirth?: string;
  gender?: string;
  address?: string;
  emergencyContact?: string;
  emergencyPhone?: string;
  bloodGroup?: string;
  allergies?: string;
  diagnosis?: string;
  notes?: string;
}

export const patientService = {
  getPatients: async (params: GetPatientsParams, role: 'doctor' | 'receptionist'): Promise<PaginatedResponse<Patient>> => {
    const { data } = await api.get(`/${role}/patients`, { params });
    return data;
  },

  getPatient: async (id: number, role: 'doctor' | 'receptionist'): Promise<Patient> => {
    const { data } = await api.get(`/${role}/patients/${id}`);
    return data;
  },

  createPatient: async (patientData: Omit<Patient, 'id' | 'createdAt' | 'updatedAt' | 'createdBy' | 'updatedBy'>): Promise<Patient> => {
    const { data } = await api.post('/receptionist/patients', patientData);
    return data;
  },

  updatePatient: async (id: number, patientData: UpdatePatientData, role: 'doctor' | 'receptionist' = 'receptionist'): Promise<Patient> => {
    const method = role === 'doctor' ? 'patch' : 'put';
    const { data } = await api[method](`/${role}/patients/${id}`, patientData);
    return data;
  },

  deletePatient: async (id: number): Promise<void> => {
    await api.delete(`/receptionist/patients/${id}`);
  },
};

import { routes } from '../common/config/routes';
import { IDevice } from '../types/device';
import axiosClient from './base';

export const getOnlineDevices = async (): Promise<IDevice[]> => {
  const url = `${routes.devices}`;
  const { data } = await axiosClient.get(url);
  return data;
};

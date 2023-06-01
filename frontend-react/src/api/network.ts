
import { routes } from "../common/config/routes";
import { INetwork } from "../types/network";
import axiosClient from "./base";

export const getNetworksByDeviceId = async (deviceId: string): Promise<INetwork[]> => {
  const url = `${routes.networks}/${deviceId}`;
  const { data } = await axiosClient.get(url);
  return data;
};
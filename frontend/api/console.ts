import { routes } from "../common/config/routes";
import { IConsole } from "../types/console";
import axiosClient from "./base";

export const getConsolesByDeviceId = async (deviceId: string): Promise<IConsole[]> => {
  const url = `${routes.consoles}/${deviceId}`;
  const { data } = await axiosClient.get(url);
  return data;
};

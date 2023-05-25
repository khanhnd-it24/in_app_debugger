import { useQuery } from "react-query";
import { getConsolesByDeviceId } from "../api/console";

export const useGetConsolesByDeviceId = (deviceId: string) => {
  return useQuery(['consoles', deviceId], () => getConsolesByDeviceId(deviceId));
}
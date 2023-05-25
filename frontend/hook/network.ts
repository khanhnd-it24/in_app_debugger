import { useQuery } from "react-query";
import { getNetworksByDeviceId } from "../api/network";

export const useGetNetworksByDeviceId = (deviceId: string) => {
  return useQuery(['networks', deviceId], () => getNetworksByDeviceId(deviceId));
}
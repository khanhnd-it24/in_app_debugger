import { useQuery } from "react-query";
import { getOnlineDevices } from "../api/device";

export const useGetOnlineDevices = (refresh: number) => {
  return useQuery(['devices', refresh], () => getOnlineDevices());
}

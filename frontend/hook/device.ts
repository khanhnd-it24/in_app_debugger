import { useQuery } from "react-query";
import { getOnlineDevices } from "../api/device";

export const useGetOnlineDevices = () => {
  return useQuery(['devices'], () => getOnlineDevices());
}

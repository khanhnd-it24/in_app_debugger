import axios from "axios";
import appConfig from "../common/config/app";

const axiosClient = axios.create({
  baseURL: `${appConfig.apiUrl}/`,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  withCredentials: true,
});

export default axiosClient;
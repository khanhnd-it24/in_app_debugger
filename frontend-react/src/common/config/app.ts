const appConfig = {
  appUrl: import.meta.env.VITE_APP_URL || "http://localhost:3000",
  apiUrl: import.meta.env.VITE_API_URL || "http://localhost:8000/apis/v1",
};

export default appConfig;
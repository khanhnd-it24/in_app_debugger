const appConfig = {
  appUrl: process.env.NEXT_PUBLIC_APP_URL || "http://localhost:3000",
  apiUrl: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000/apis/v1",
};

export default appConfig;
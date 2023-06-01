const mqttConfig = {
  protocol: import.meta.env.VITE_MQTT_PROTOCOL || "ws",
  host: import.meta.env.VITE_MQTT_HOST || "localhost:9001/mqtt",
  username: import.meta.env.VITE_MQTT_USERNAME || "admin",
  password: import.meta.env.VITE_MQTT_PASSWORD || "password",
};

export default mqttConfig;
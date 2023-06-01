import  mqtt  from 'mqtt';
import { MqttClient } from 'mqtt';
import mqttConfig from '../../common/config/mqtt';

const mqttUri = `${mqttConfig.protocol}://${mqttConfig.host}`;

const options = (clientId: string) => {
  return {
    username: mqttConfig.username,
    password: mqttConfig.password,
    clientId: clientId,
  };
};

export const getClient = (clientId: string): MqttClient => {
  const client = mqtt.connect(mqttUri, options(clientId))
  return client
}

import { Tabs } from 'antd';
import { useEffect, useState } from 'react';
import { DeviceTab } from '../components/DeviceTab';
import { Loading } from '../components/Loading';
import { useGetOnlineDevices } from '../hook/device';
import { getClient } from '../utils/transports/mqtt';
import { v4 } from 'uuid';

const topic = "devices"

const HomeComponent = () => {
  const clientId = v4()
  const [refreshKey, setRefreshKey] = useState(0);
  const { data: devices, isLoading } = useGetOnlineDevices(refreshKey);

  useEffect(() => {
    console.log('Connecting to MQTT broker');

    const client = getClient(clientId)

    client.on('connect', () => {
      console.log('Connected to MQTT broker');
    });

    client.on('error', function (error) {
      console.log('error', error)
    })

    client.on('message', (_topic, message) => {
      console.log(message.toString());

      if (_topic === topic) {
        setRefreshKey(prev => prev + 1);
      }
    });

    client.subscribe(topic, () => {
      console.log(`Subscribe successfully!`);
    });

    return () => {
      client.end();
    };
  }, [clientId]);

  if (isLoading) return <Loading />

  return (
    <Tabs
      defaultActiveKey="1"
      tabPosition={"left"}
      items={(devices ?? []).map((device, i) => {
        return {
          label: device.deviceName,
          key: `item${i}`,
          children: <DeviceTab device={device} />,
        };
      })}
    />
  )
}

export default HomeComponent
import { Col, Row } from "antd";
import { useEffect, useState } from "react";
import { useGetNetworksByDeviceId } from "../hook/network";
import { INetwork } from "../types/network";
import { getClient } from '../utils/transports/mqtt';
import { Loading } from "./Loading";
import NetworkItemView from "./NetworkItemView";
import NetworkList from "./NetworkList";
import './network.css';
import { v4 } from "uuid";

export const NetworkComponent = ({ deviceId }: { deviceId: string }) => {
  const clientId = v4()
  const [activeItem, setActiveItem] = useState<INetwork | null>(null);
  const [networks, setNetworks] = useState<INetwork[]>([])
  const { data, isLoading } = useGetNetworksByDeviceId(deviceId);

  useEffect(() => {
    if (data != null) {
      setNetworks(data)
    }
  }, [data])

  useEffect(() => {
    const topic = `networks/${deviceId}`

    const client = getClient(clientId)

    client.on('connect', () => {
      console.log('Connected to MQTT broker network');
    });

    client.on('message', (_topic, message) => {

      if (_topic === topic) {
        const network: INetwork = JSON.parse(message.toString())
        setNetworks(prev => [...prev, network])
      }
    });

    client.subscribe(topic, () => {
      console.log(`Subscribe successfully! ${topic}`);
    });

    return () => {
      client.end();
    };
  }, [clientId, deviceId]);

  if (isLoading) return <Loading />
  return (
    <Row style={{ width: "100%" }}>
      <Col span={12}>
        <NetworkList networks={networks} activeItem={activeItem} setActiveItem={setActiveItem} />
      </Col>
      <Col span={1} />
      <Col span={11}>
        {activeItem !== null && <NetworkItemView network={activeItem} />}
      </Col>
    </Row>
  )
}
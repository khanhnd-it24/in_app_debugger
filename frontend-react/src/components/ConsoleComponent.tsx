import VirtualList from 'rc-virtual-list';
import { useEffect, useState } from 'react';
import { v4 } from 'uuid';
import { useGetConsolesByDeviceId } from '../hook/console';
import { IConsole } from '../types/console';
import { getClient } from '../utils/transports/mqtt';
import { Loading } from './Loading';

export const ConsoleComponent = ({ deviceId }: { deviceId: string }) => {
  const clientId = v4()
  const [consoles, setConsoles] = useState<IConsole[]>([])
  const { data, isLoading } = useGetConsolesByDeviceId(deviceId);

  useEffect(() => {
    if (data != null) {
      setConsoles(data)
    }
  }, [data])

  useEffect(() => {
    const topic = `consoles/${deviceId}`

    const client = getClient(clientId)

    client.on('connect', () => {
      console.log('Connected to MQTT broker console');
    });

    client.on('message', (_topic, message) => {

      if (_topic === topic) {
        const _console: IConsole = JSON.parse(message.toString())
        console.log(_console);
        setConsoles(prev => [...prev, _console])
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
    <VirtualList
      data={consoles ?? []}
      height={700}
      itemHeight={consoles?.length ?? 0}
      itemKey={(item: IConsole) => `Console${item._id}`}
    >
      {(item: IConsole) => (
        <div key={item._id} style={{ display: "flex" }} >
          <p style={{ margin: 4, textAlign: "left" }} dangerouslySetInnerHTML={{ __html: item.content.replace(/\n/g, '<br>') }}></p>
        </div>
      )}
    </VirtualList>
  )
}

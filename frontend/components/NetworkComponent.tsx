import { Button, List } from "antd";
import { useGetNetworksByDeviceId } from "../hook/network";
import { Loading } from "./Loading";
import VirtualList from 'rc-virtual-list';
import { INetwork } from "../types/network";

export const NetworkComponent = ({ deviceId }: { deviceId: string }) => {
  const { data: networks, isLoading } = useGetNetworksByDeviceId(deviceId);

  if (isLoading) return <Loading />
  return (
    <List>
      <VirtualList
        data={networks ?? []}
        height={400}
        itemHeight={networks?.length ?? 0}
        itemKey="console"
      >
        {(item: INetwork) => (
          <List.Item key={item._id}>
            <List.Item.Meta
              title={<>{item.path}</>}
              description={item.method}
            />
            {item.statusCode < 300 ?
              <Button type="primary" shape="round" size="small">
                {`${item.statusCode} OK`}
              </Button>
              : <Button type="primary" shape="round" size="small" danger>
                {`${item.statusCode} ERROR`}
              </Button>
            }
          </List.Item>
        )}
      </VirtualList>
    </List>
  )
}
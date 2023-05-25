import { List } from 'antd'
import React from 'react'
import VirtualList from 'rc-virtual-list';
import { Loading } from './Loading';
import { useGetConsolesByDeviceId } from '../hook/console';
import { IConsole } from '../types/console';
import dayjs from 'dayjs';

export const ConsoleComponent = ({ deviceId }: { deviceId: string }) => {
  const { data: consoles, isLoading } = useGetConsolesByDeviceId(deviceId);

  if (isLoading) return <Loading />
  return (
    <List>
      <VirtualList
        data={consoles ?? []}
        height={400}
        itemHeight={consoles?.length ?? 0}
        itemKey="console"
      >
        {(item: IConsole) => (
          <List.Item key={item._id}>
            <div>{item.content}</div>
            <div>{dayjs(item.createdAt).format("YYYY-MM-DD HH:mm:ss")}</div>
          </List.Item>
        )}
      </VirtualList>
    </List>
  )
}

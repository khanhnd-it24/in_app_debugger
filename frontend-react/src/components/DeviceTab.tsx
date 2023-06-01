import { IDevice } from '../types/device'
import { Tabs } from 'antd'
import { ConsoleComponent } from './ConsoleComponent'
import { NetworkComponent } from './NetworkComponent'

interface DeviceTabProps {
  device: IDevice
}

export const DeviceTab = ({ device }: DeviceTabProps) => {
  return (
    <Tabs
      key={`DeviceTab ${device.deviceId}`}
      defaultActiveKey="1"
      tabPosition={"top"}
      items={[
        {
          label: "Console",
          key: `Console ${device.deviceId}`,
          children: <ConsoleComponent deviceId={device.deviceId} />,
        },
        {
          label: "Network",
          key: `Network ${device.deviceId}`,
          children: <NetworkComponent deviceId={device.deviceId} />,
        },
      ]}
    />
  )
}

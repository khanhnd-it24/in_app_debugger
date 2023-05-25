import type { NextPage } from 'next'
import { Tabs } from 'antd'
import { DeviceTab } from '../components/DeviceTab'
import { useGetOnlineDevices } from '../hook/device'
import { Loading } from '../components/Loading'

const Home: NextPage = () => {
  const { data: devices, isLoading } = useGetOnlineDevices();

  if (isLoading) return <Loading />
  return (
    <div>
      <h1 style={{margin: 16}}>In App Debugger</h1>
      <Tabs
        defaultActiveKey="1"
        tabPosition={"left"}
        items={(devices??[]).map((device, i) => {
          return {
            label: device.deviceName,
            key: `item${i}`,
            children: <DeviceTab device={device} />,
          };
        })}
      />
    </div>

  )
}

export default Home

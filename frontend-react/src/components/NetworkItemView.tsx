import { INetwork } from '../types/network'
import { Tabs } from 'antd'
import ViewJson from './ViewJson'

const NetworkItemView = ({ network }: { network: INetwork }) => {
  return (
    <Tabs
      key={`DeviceTab ${network._id}`}
      defaultActiveKey="1"
      tabPosition={"top"}
      items={[
        {
          label: "Request",
          key: `Request ${network._id}`,
          children: <ViewJson json={JSON.parse(network.request)}/>,
        },
        {
          label: "Response",
          key: `Response ${network._id}`,
          children: <ViewJson json={JSON.parse(network.response)}/>,
        },
      ]}
    />
  )
}

export default NetworkItemView
import React from 'react'
import { IDevice } from '../types/device'
import { Col, Row, Typography } from 'antd'
import { ConsoleComponent } from './ConsoleComponent'
import { NetworkComponent } from './NetworkComponent'

const { Title } = Typography;

interface DeviceTabProps {
  device: IDevice
}

export const DeviceTab = ({ device }: DeviceTabProps) => {
  return (
    <Row key={`device${device.deviceId}`} style={{ paddingRight: 16 }}>
      <Col span={11}>
        <Title >Consoles</Title>
        <ConsoleComponent deviceId={device.deviceId} />
      </Col>
      <Col span={2} />
      <Col span={11}>
        <Title>Networks</Title>
        <NetworkComponent deviceId={device.deviceId} />
      </Col>
    </Row>
  )
}

import { Spin } from 'antd'
import React from 'react'

export const Loading = () => {
  return (
    <Spin tip="Loading" size="large">
      <div className="content" />
    </Spin>
  )
}

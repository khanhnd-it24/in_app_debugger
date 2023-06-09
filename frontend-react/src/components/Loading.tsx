import { Spin } from 'antd'

export const Loading = () => {
  return (
    <Spin tip="Loading" size="large">
      <div className="content" />
    </Spin>
  )
}

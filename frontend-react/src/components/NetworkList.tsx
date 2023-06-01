import VirtualList from 'rc-virtual-list'
import { Button, Col, List, Row } from 'antd'
import { INetwork } from '../types/network'

const NetworkList = ({ networks, activeItem, setActiveItem }:
  {
    networks: INetwork[],
    activeItem: INetwork | null,
    setActiveItem: (activeItem: INetwork | null) => void
  }) => {
  const handleItemClick = (item: INetwork) => {
    setActiveItem(item);
  };

  return (
    <List>
      <VirtualList
        data={networks ?? []}
        height={700}
        itemHeight={networks?.length ?? 0}
        itemKey={(item: INetwork) => `Network${item._id}`}
      >
        {(item: INetwork) => (
          <List.Item
            key={item._id}
            className={activeItem?._id === item._id ? 'active' : ''}
            onClick={() => handleItemClick(item)}
          >
            <Row style={{ width: "100%" }} justify={"start"}>
              <Col span={2}>
                <h3 style={{ margin: 0 }}>{item.method}</h3>
              </Col>
              <Col span={20}>
                <h5 style={{ margin: 0, textAlign: "left", paddingTop: 4 }}>{item.path}</h5>
              </Col>
              <Col span={2}>
                {item.statusCode < 300 ?
                  <Button type="primary" shape="round" size="small">
                    {`${item.statusCode}`}
                  </Button>
                  : <Button type="primary" shape="round" size="small" danger>
                    {`${item.statusCode}`}
                  </Button>
                }
              </Col>
            </Row>
          </List.Item>
        )}
      </VirtualList>
    </List>
  )
}

export default NetworkList
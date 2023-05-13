package domains

import (
	"backend/src/common"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Network struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId  string             `bson:"device_id"`
	Request   string             `bson:"request"`
	Response  string             `bson:"response"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewNetwork() *Network {
	return &Network{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (n *Network) SetDeviceId(deviceId string) *Network {
	n.DeviceId = deviceId
	return n
}

func (n *Network) SetRequest(request string) *Network {
	n.Request = request
	return n
}

func (n *Network) SetResponse(response string) *Network {
	n.Response = response
	return n
}

type NetworkRepo interface {
	Create(ctx context.Context, network *Network) (*Network, *common.Error)
	GetNetworksByDeviceId(ctx context.Context, deviceId string) ([]*Network, *common.Error)
}

func (n *Network) CollectionName() string {
	return "network"
}

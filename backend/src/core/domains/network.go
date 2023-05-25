package domains

import (
	"backend/src/common"
	"backend/src/common/enums"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Network struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	DeviceId   string             `bson:"device_id" json:"deviceId"`
	Method     enums.HttpMethod   `bson:"method" json:"method"`
	Path       string             `bson:"path" json:"path"`
	StatusCode uint               `bson:"status_code" json:"statusCode"`
	Request    string             `bson:"request" json:"request"`
	Response   string             `bson:"response" json:"response"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updatedAt"`
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

func (n *Network) SetMethod(method enums.HttpMethod) *Network {
	n.Method = method
	return n
}

func (n *Network) SetPath(path string) *Network {
	n.Path = path
	return n
}

func (n *Network) SetStatusCode(statusCode uint) *Network {
	n.StatusCode = statusCode
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

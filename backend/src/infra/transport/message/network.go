package message

import (
	"backend/src/common/enums"
	"backend/src/core/domains"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NetworkMsg struct {
	Id         primitive.ObjectID `json:"_id"`
	DeviceId   string             `json:"deviceId"`
	Method     enums.HttpMethod   `json:"method"`
	Path       string             `json:"path"`
	StatusCode uint               `json:"statusCode"`
	Request    string             `json:"request"`
	Response   string             `json:"response"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
}

func (i *NetworkMsg) Topic() string {
	return fmt.Sprintf(MqttNetwork, i.DeviceId)
}

func (i *NetworkMsg) Payload() string {
	val, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(val)
}

func NewNetworkMsg(network *domains.Network) *NetworkMsg {
	return &NetworkMsg{
		Id:         network.Id,
		DeviceId:   network.DeviceId,
		Method:     network.Method,
		Path:       network.Path,
		StatusCode: network.StatusCode,
		Request:    network.Request,
		Response:   network.Response,
		CreatedAt:  network.CreatedAt,
		UpdatedAt:  network.UpdatedAt,
	}
}

package message

import (
	"backend/src/core/domains"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConsoleMsg struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	DeviceId  string             `bson:"device_id" json:"deviceId"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
}

func (i *ConsoleMsg) Topic() string {
	return fmt.Sprintf(MqttConsole, i.DeviceId)
}

func (i *ConsoleMsg) Payload() string {
	val, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(val)
}

func NewConsoleMsg(console *domains.Console) *ConsoleMsg {
	return &ConsoleMsg{
		Id:        console.Id,
		DeviceId:  console.DeviceId,
		Content:   console.Content,
		CreatedAt: console.CreatedAt,
		UpdatedAt: console.UpdatedAt,
	}
}

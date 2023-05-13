package domains

import (
	"backend/src/common"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Console struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId  string             `bson:"device_id"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewConsole() *Console {
	return &Console{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Console) SetDeviceId(deviceId string) *Console {
	c.DeviceId = deviceId
	return c
}

func (c *Console) SetContent(content string) *Console {
	c.Content = content
	return c
}

type ConsoleRepo interface {
	Create(ctx context.Context, console *Console) (*Console, *common.Error)
	GetConsolesByDeviceId(ctx context.Context, deviceId string) ([]*Console, *common.Error)
}

func (c *Console) CollectionName() string {
	return "console"
}

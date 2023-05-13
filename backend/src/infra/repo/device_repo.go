package repo

import (
	"backend/src/common"
	"backend/src/core/domains"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/appengine/log"
)

func NewDeviceRepo(db *mongo.Database) domains.DeviceRepo {
	return &DeviceRepo{
		collection: db.Collection(new(domains.Device).CollectionName()),
	}
}

type DeviceRepo struct {
	collection *mongo.Collection
}

func (d DeviceRepo) Upsert(ctx context.Context, device *domains.Device) (*domains.Device, *common.Error) {
	filter := bson.D{{Key: "device_id", Value: device.DeviceId}}
	update := bson.D{{Key: "$set", Value: device}}
	opts := options.Update().SetUpsert(true)
	_, err := d.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	return device, nil
}

func (d DeviceRepo) GetDeviceByDeviceId(ctx context.Context, deviceId string) (*domains.Device, *common.Error) {
	invoice := new(domains.Device)
	filter := bson.D{{Key: "device_id", Value: deviceId}}
	err := d.collection.FindOne(ctx, filter).Decode(&invoice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrNotFound(ctx, "device", "not found")
		}
		return nil, common.ErrSystemError(ctx, err.Error())
	}
	return invoice, nil
}

func (d DeviceRepo) GetAllDevices(ctx context.Context) ([]*domains.Device, *common.Error) {
	devices := make([]*domains.Device, 0)
	filter := bson.D{{}}
	cursor, err := d.collection.Find(ctx, filter)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	for cursor.Next(ctx) {
		var device *domains.Device
		err = cursor.Decode(&device)
		if err != nil {
			log.Errorf(ctx, "Cannot decode device %+v", err)
		} else {
			devices = append(devices, device)
		}
	}
	return devices, nil
}

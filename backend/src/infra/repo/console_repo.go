package repo

import (
	"backend/src/common"
	"backend/src/core/domains"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/appengine/log"
)

func NewConsoleRepo(db *mongo.Database) domains.ConsoleRepo {
	return &ConsoleRepo{
		collection: db.Collection(new(domains.Network).CollectionName()),
	}
}

type ConsoleRepo struct {
	collection *mongo.Collection
}

func (c ConsoleRepo) Create(ctx context.Context, console *domains.Console) (*domains.Console, *common.Error) {
	_, err := c.collection.InsertOne(ctx, console)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	return console, nil
}

func (c ConsoleRepo) GetConsolesByDeviceId(ctx context.Context, deviceId string) ([]*domains.Console, *common.Error) {
	consoles := make([]*domains.Console, 0)
	filter := bson.D{{Key: "device_id", Value: deviceId}}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	for cursor.Next(ctx) {
		var console *domains.Console
		err = cursor.Decode(&console)
		if err != nil {
			log.Errorf(ctx, "Cannot decode console %+v", err)
		} else {
			consoles = append(consoles, console)
		}
	}

	return consoles, nil
}

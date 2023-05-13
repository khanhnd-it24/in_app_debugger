package repo

import (
	"backend/src/common"
	"backend/src/core/domains"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/appengine/log"
)

func NewNetworkRepo(db *mongo.Database) domains.NetworkRepo {
	return &NetworkRepo{
		collection: db.Collection(new(domains.Network).CollectionName()),
	}
}

type NetworkRepo struct {
	collection *mongo.Collection
}

func (n NetworkRepo) Create(ctx context.Context, network *domains.Network) (*domains.Network, *common.Error) {
	_, err := n.collection.InsertOne(ctx, network)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	return network, nil
}

func (n NetworkRepo) GetNetworksByDeviceId(ctx context.Context, deviceId string) ([]*domains.Network, *common.Error) {
	networks := make([]*domains.Network, 0)
	filter := bson.D{{Key: "device_id", Value: deviceId}}

	cursor, err := n.collection.Find(ctx, filter)
	if err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}

	for cursor.Next(ctx) {
		var network *domains.Network
		err = cursor.Decode(&network)
		if err != nil {
			log.Errorf(ctx, "Cannot decode network %+v", err)
		} else {
			networks = append(networks, network)
		}
	}

	return networks, nil
}

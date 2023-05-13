package bootstrap

import (
	"backend/src/common/configs"
	"backend/src/common/log"
	"backend/src/infra/repo"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

func BuildStorageModules() fx.Option {
	return fx.Options(
		fx.Provide(newMongoDB),
		fx.Provide(newCacheRedis),

		fx.Provide(repo.NewDeviceRepo),
		fx.Provide(repo.NewNetworkRepo),
		fx.Provide(repo.NewConsoleRepo),
	)
}

func newMongoDB(lc fx.Lifecycle, logger *zap.SugaredLogger) *mongo.Database {
	logger.Debugf("Coming Create Storage")
	cf := configs.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.ClientOptions{}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cf.Mongo.Uri), &opts)
	if err != nil {
		logger.Fatalf("connect mongo db error:[%s]", err.Error())
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		logger.Fatalf("ping mongo db error:[%s]", err.Error())
	}
	db := client.Database(cf.Mongo.DB)
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Info("Coming OnStop Storage")
			return client.Disconnect(ctx)
		},
	})
	return db
}

func newCacheRedis() redis.UniversalClient {
	cf := configs.Get().Redis
	hosts := cf.Hosts
	var client redis.UniversalClient
	isClusterMode := len(hosts) > 1
	if isClusterMode {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    hosts,
			Username: cf.Username,
			Password: cf.Password,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     hosts[0],
			Username: cf.Username,
			Password: cf.Password,
		})
	}

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.GetLogger().GetZap().Fatalf("ping redis error, err:[%s]", err.Error())
	}
	return client
}

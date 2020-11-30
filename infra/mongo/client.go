package mongo

import (
	"context"
	"gimServer/conf"
	"gimServer/infra/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(config *conf.Config) *mongo.Client {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(config.Mongo.Url).
			SetMaxPoolSize(config.Mongo.MaxPoolSize).
			SetMinPoolSize(config.Mongo.MinPoolSize))
	utils.Must(err)
	err = client.Connect(context.TODO())
	utils.Must(err)
	return client
}

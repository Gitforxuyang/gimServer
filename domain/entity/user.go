package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Uid        int64              `bson:"uid"`
	Token      string             `bson:"token"`
	Uuid       int64              `bson:"uuid"`
	SdkVersion string             `bson:"sdkVersion"`
	DeviceId   string             `bson:"deviceId"`
	Platform   string             `bson:"platform"`
	Model      string             `bson:"model"`
	System     string             `bson:"system"`
}



package repo

import (
	"context"
	"gimServer/domain/entity"
	"gimServer/infra/err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *domainRepo) SelectUserByToken(ctx context.Context, uid int64, token string) (*entity.User, error) {
	r := m.user.FindOne(ctx, bson.M{"uid": uid, "token": token})
	if r.Err() == mongo.ErrNoDocuments {
		return nil, err.UserNotFoundError
	}
	if err := r.Err(); err != nil {
		return nil, err
	}
	user := &entity.User{}
	err := r.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *domainRepo) UpdateUserById(ctx context.Context, uid int64, uuid int64) error {
	_, err := m.user.UpdateOne(ctx, bson.M{"uid": uid}, bson.M{"$set": bson.M{"uuid": uuid}})
	if err != nil {
		return err
	}
	return nil
}

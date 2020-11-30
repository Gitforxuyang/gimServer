package repo

import (
	"context"
	"gimServer/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *domainRepo) SaveMsg(ctx context.Context, msg *entity.Msg) error {
	_, err := m.msg.InsertOne(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *domainRepo) SaveInbox(ctx context.Context, msg *entity.InboxMsg) error {
	_, err := m.inbox.InsertOne(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *domainRepo) SelectMsgBySeq(ctx context.Context, seq int64) (*entity.Msg, error) {
	r := m.msg.FindOne(ctx, bson.M{"seq": seq})
	if e := r.Err(); e != nil {
		if e == mongo.ErrNoDocuments {
			return &entity.Msg{}, nil
		}
		return nil, e
	}
	msg := entity.Msg{}
	err := r.Decode(&msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

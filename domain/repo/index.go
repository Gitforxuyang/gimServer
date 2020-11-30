package repo

import (
	"context"
	"gimServer/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDomainRepo interface {
	SelectUserByToken(ctx context.Context, uid int64, token string) (*entity.User, error)
	UpdateUserById(ctx context.Context, uid int64, uuid int64) error
	SaveMsg(ctx context.Context, msg *entity.Msg) error
	SelectMsgBySeq(ctx context.Context, seq int64) (*entity.Msg, error)
	SaveInbox(ctx context.Context, msg *entity.InboxMsg) error
}

func NewDomainRepo(mongoClient *mongo.Client) IDomainRepo {
	repo := domainRepo{}
	db:=mongoClient.Database("im")
	repo.user = db.Collection("user")
	repo.msg = db.Collection("msg")
	repo.inbox = db.Collection("inbox")
	return &repo
}

type domainRepo struct {
	user  *mongo.Collection
	msg   *mongo.Collection
	inbox *mongo.Collection
}

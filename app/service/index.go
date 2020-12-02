package service

import (
	"context"
	"gimServer/domain/entity"
	"gimServer/domain/repo"
	"gimServer/infra/rabbitmq"
	"github.com/go-redis/redis/v7"
)

type IService interface {
	Auth(ctx context.Context, uid, uuid int64, token, sdkVersion, deviceId, platform, model, system string) error
	SendMsg(ctx context.Context, _type, action int32, seq, from, to int64, content string) (*entity.Msg, error)
}

type service struct {
	repo  repo.IDomainRepo
	mq    *rabbitmq.Queue
	redis *redis.Client
}

func NewService(repo repo.IDomainRepo, mq *rabbitmq.Queue, redis *redis.Client) IService {
	return &service{repo: repo, mq: mq, redis: redis}
}

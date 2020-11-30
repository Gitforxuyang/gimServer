package service

import (
	"context"
	"gimServer/domain/entity"
	"gimServer/domain/repo"
)

type IService interface {
	Auth(ctx context.Context, uid, uuid int64, token, sdkVersion, deviceId, platform, model, system string) error
	SendMsg(ctx context.Context, _type, action int32, seq, from, to int64, content string) (*entity.Msg, error)
}

type service struct {
	repo repo.IDomainRepo

}

func NewService(repo repo.IDomainRepo) IService {
	return &service{repo: repo}
}

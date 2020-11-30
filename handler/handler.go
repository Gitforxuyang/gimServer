package handler

import (
	"context"
	"gimServer/app/service"
	"gimServer/infra/err"
	"gimServer/proto"
)

type Handler struct {
	svc service.IService
}

func NewHandler(svc service.IService) *Handler {
	return &Handler{svc}
}
func (m *Handler) Ping(ctx context.Context, req *im.Nil) (*im.Nil, error) {
	return &im.Nil{}, nil
}

func (m *Handler) Auth(ctx context.Context, req *im.AuthReq) (*im.Nil, error) {
	if req.Uid == 0 || req.Token == "" || req.Uuid == 0 {
		return nil, err.ParamsError
	}
	err := m.svc.Auth(ctx, req.Uid, req.Uuid, req.Token, req.SdkVersion, req.DeviceId, req.Platform, req.Model, req.System)
	if err != nil {
		return nil, err
	}
	return &im.Nil{}, nil
}

func (m *Handler) SendMsg(ctx context.Context, req *im.SendMsgReq) (*im.SendMsgResp, error) {
	if req.Type == 0 || req.To == 0 || req.From == 0 || req.Seq == 0 || req.Action == 0 || req.Content == "" {
		return nil, err.ParamsError
	}
	msg, err := m.svc.SendMsg(ctx, req.Type, int32(req.Action), req.Seq, req.From, req.To, req.Content)
	if err != nil {
		return nil, err
	}
	return &im.SendMsgResp{Seq: msg.Seq, MsgId: msg.MsgId}, nil
}

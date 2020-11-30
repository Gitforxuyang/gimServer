package service

import (
	"context"
	"fmt"
	"gimServer/domain/entity"
	"gimServer/infra/utils"
)

func (m *service) SendMsg(ctx context.Context, _type, action int32, seq, from, to int64, content string) (*entity.Msg, error) {
	msg, err := m.repo.SelectMsgBySeq(ctx, seq)
	if err != nil {
		return nil, err
	}
	//如果消息已存在，则直接返回
	if msg.MsgId != 0 {
		return msg, nil
	}
	msg = &entity.Msg{}
	msg.MsgId = utils.GetSnowflakeId()
	msg.Content = content
	msg.Action = action
	msg.Seq = seq
	msg.From = from
	msg.To = to
	msg.Type = _type
	msg.Channel = _getChannel(_type, from, to)
	err = m.repo.SaveMsg(ctx, msg)
	if err != nil {
		return nil, err
	}
	err = m._saveInbox(ctx, _type, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *service) _saveInbox(ctx context.Context, _type int32, msg *entity.Msg) error {
	if _type == 1 {
		return m.repo.SaveInbox(ctx, &entity.InboxMsg{Uid: msg.To, SeqId: utils.GetSnowflakeId(), MsgId: msg.MsgId})
	}
	//如果是群聊
	if _type == 2 {

	}
	return nil
}
func _getChannel(_type int32, from, to int64) string {
	//如果是单聊
	if _type == 1 {
		if from > to {
			return fmt.Sprintf("s:%d:%d", to, from)
		} else {
			return fmt.Sprintf("s:%d:%d", from, to)
		}
	}
	//如果是群聊
	if _type == 2 {
		return fmt.Sprintf("t:%d", to)
	}
	if _type == 3 {
		return fmt.Sprintf("c:%d", to)
	}
	return ""
}

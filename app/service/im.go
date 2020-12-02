package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gimServer/domain/entity"
	"gimServer/infra/utils"
	"strconv"
)

const (
	//基本功能
	CmdId_Ping uint8 = 1
	CmdId_Pong uint8 = 2

	//基础项
	CmdId_AuthReq     uint8 = 21 //认证请求
	CmdId_AuthResp    uint8 = 22 //认证返回
	CmdId_LogoutReq   uint8 = 23 //退出
	CmdId_LoutoutResp uint8 = 24
	CmdId_KickOut     uint8 = 26 //踢出

	//偏业务逻辑项
	CmdId_SendMessageReq   uint8 = 101 //发送消息
	CmdId_SendMessageResp  uint8 = 102 //发送消息
	CmdId_Notify           uint8 = 104
	CmdId_NotifyAck        uint8 = 103
	CmdId_SyncMessageReq   uint8 = 105
	CmdId_SyncMessageResp  uint8 = 106
	CmdId_FetchMessageReq  uint8 = 107
	CmdId_FetchMessageResp uint8 = 108
	CmdId_SyncLastIdReq    uint8 = 109
	CmdId_SyncLastIdResp   uint8 = 110
)

func (m *service) SendMsg(ctx context.Context, _type, action int32, seq, from, to int64, content string) (*entity.Msg, error) {
	msg, err := m.repo.SelectMsgBySeq(ctx, seq)
	if err != nil {
		return nil, err
	}
	//如果消息已存在，则直接返回
	if msg.MsgId != 0 {
		//TODO:为了测试发送通知
		err = m._sendNotify(_type, CmdId_SendMessageResp, msg.MsgId, msg.To)
		if err != nil {
			return nil, err
		}
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
	//发送通知
	err = m._sendNotify(_type, CmdId_SendMessageResp, msg.MsgId, msg.To)
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

type QueueHeader struct {
	Type  int32 `json:"type"`  //消息类别  单聊 群聊 聊天室
	CmdId uint8 `json:"cmdId"` //操作码id
	MsgId int64 `json:"msgId"` //消息id
	To    int64 `json:"to"`    //发送目标
	UUID  int64 `json:"uuid"`  //如果指定连接id则不为0
}

func (m *service) _sendNotify(_type int32, cmdId uint8, msgId, uid int64) error {
	routingKey := ""
	//如果是单聊，则找到接收方的node消息
	notify := QueueHeader{Type: _type, CmdId: cmdId, MsgId: msgId, To: uid}
	if _type == 1 {
		node, uuid, err := m._getUserNode(uid)
		if err != nil {
			return err
		}
		//对方不在线，则忽略通知
		if node == "" {
			return nil
		}
		routingKey = fmt.Sprintf("gim_node.%s", node)
		notify.UUID = uuid
	}
	body, err := json.Marshal(&notify)
	if err != nil {
		return err
	}
	headers, err := utils.StructToMap(notify)
	if err != nil {
		return err
	}
	return m.mq.Publish(routingKey, headers, body)
}

func (m *service) _getUserNode(uid int64) (string, int64, error) {
	fields, err := m.redis.HGetAll(fmt.Sprintf("u:%d", uid)).Result()
	if err != nil {
		return "", 0, err
	}
	node := fields["node"]
	uuid, _ := strconv.ParseInt(fields["uuid"], 10, 64)
	return node, uuid, nil
}

package entity

type Msg struct {
	Channel string `bson:"channel"` //信道 对于单聊是 s:uid:uid 对于群聊是 t:tid
	MsgId   int64  `bson:"msgId"`
	Seq     int64  `bson:"seq"`  //客户端发送id
	Type    int32  `bson:"type"` //  1-单聊 2-群聊 3-聊天室
	Action  int32  `bson:"action"`
	From    int64  `bson:"from"`
	To      int64  `bson:"to"`
	Content string `bson:"content"`
}

type InboxMsg struct {
	Uid   int64 `bson:"uid"`
	SeqId int64 `bson:"seqId"` //最好是设置为数据库级的自增id
	MsgId int64 `bson:"msgId"`
}

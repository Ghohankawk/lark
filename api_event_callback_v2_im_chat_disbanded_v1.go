package lark

import (
	"context"
)

// EventV2IMChatDisbandedV1
//
// 群组被解散后触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [解散群] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/disbanded
func (r *EventCallbackAPI) HandlerEventV2IMChatDisbandedV1(f eventV2IMChatDisbandedV1Handler) {
	r.cli.eventHandler.eventV2IMChatDisbandedV1Handler = f
}

type eventV2IMChatDisbandedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatDisbandedV1) (string, error)

type EventV2IMChatDisbandedV1 struct {
	ChatID     string                              `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventV2IMChatDisbandedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
}

type EventV2IMChatDisbandedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

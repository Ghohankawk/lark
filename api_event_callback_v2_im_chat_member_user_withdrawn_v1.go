package lark

import (
	"context"
)

// EventV2IMChatMemberUserWithdrawnV1
//
// 撤销拉用户进群后触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [撤销拉用户进群] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/withdrawn
func (r *EventCallbackAPI) HandlerEventV2IMChatMemberUserWithdrawnV1(f eventV2IMChatMemberUserWithdrawnV1Handler) {
	r.cli.eventHandler.eventV2IMChatMemberUserWithdrawnV1Handler = f
}

type eventV2IMChatMemberUserWithdrawnV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatMemberUserWithdrawnV1) (string, error)

type EventV2IMChatMemberUserWithdrawnV1 struct {
	ChatID     string                                        `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventV2IMChatMemberUserWithdrawnV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
	Users      []*EventV2IMChatMemberUserWithdrawnV1User     `json:"users,omitempty"`       // 被撤销加群的用户列表
}

type EventV2IMChatMemberUserWithdrawnV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2IMChatMemberUserWithdrawnV1User struct {
	Name   string                                        `json:"name,omitempty"`    // 用户名字
	UserID *EventV2IMChatMemberUserWithdrawnV1UserUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventV2IMChatMemberUserWithdrawnV1UserUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

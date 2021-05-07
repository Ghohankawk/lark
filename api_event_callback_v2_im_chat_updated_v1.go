package lark

import (
	"context"
)

// EventV2IMChatUpdatedV1
//
// 群组配置被修改后触发此事件，包含：
// - 群主转移
// - 群基本信息修改(群头像/群名称/群描述/群国际化名称)
// - 群权限修改(加人入群权限/群编辑权限/at所有人权限/群分享权限)
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [群配置修改] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/updated
func (r *EventCallbackAPI) HandlerEventV2IMChatUpdatedV1(f eventV2IMChatUpdatedV1Handler) {
	r.cli.eventHandler.eventV2IMChatUpdatedV1Handler = f
}

type eventV2IMChatUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatUpdatedV1) (string, error)

type EventV2IMChatUpdatedV1 struct {
	ChatID        string                               `json:"chat_id,omitempty"`        // 群组 ID
	OperatorID    *EventV2IMChatUpdatedV1OperatorID    `json:"operator_id,omitempty"`    // 用户 ID
	AfterChange   *EventV2IMChatUpdatedV1AfterChange   `json:"after_change,omitempty"`   // 更新后的群信息
	BeforeChange  *EventV2IMChatUpdatedV1BeforeChange  `json:"before_change,omitempty"`  // 更新前的群信息
	ModeratorList *EventV2IMChatUpdatedV1ModeratorList `json:"moderator_list,omitempty"` // 群可发言成员名单的变更信息
}

type EventV2IMChatUpdatedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2IMChatUpdatedV1AfterChange struct {
	Avatar                 string                                    `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                    `json:"name,omitempty"`                     // 群名称
	Description            string                                    `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                                `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string                                    `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    string                                    `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        string                                    `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         string                                    `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     string                                    `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  string                                    `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string                                    `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   string                                    `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventV2IMChatUpdatedV1AfterChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

type EventV2IMChatUpdatedV1AfterChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2IMChatUpdatedV1BeforeChange struct {
	Avatar                 string                                     `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                     `json:"name,omitempty"`                     // 群名称
	Description            string                                     `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                                 `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string                                     `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    string                                     `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        string                                     `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         string                                     `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     string                                     `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  string                                     `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string                                     `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   string                                     `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventV2IMChatUpdatedV1BeforeChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

type EventV2IMChatUpdatedV1BeforeChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2IMChatUpdatedV1ModeratorList struct {
	AddedMemberList   []*EventV2IMChatUpdatedV1ModeratorListAddedMember   `json:"added_member_list,omitempty"`   // 被添加进可发言名单的用户列表（列表中一定会有owner）
	RemovedMemberList []*EventV2IMChatUpdatedV1ModeratorListRemovedMember `json:"removed_member_list,omitempty"` // 被移除出可发言名单的用户列表
}

type EventV2IMChatUpdatedV1ModeratorListAddedMember struct {
	UserID *EventV2IMChatUpdatedV1ModeratorListAddedMemberUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventV2IMChatUpdatedV1ModeratorListAddedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2IMChatUpdatedV1ModeratorListRemovedMember struct {
	UserID *EventV2IMChatUpdatedV1ModeratorListRemovedMemberUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventV2IMChatUpdatedV1ModeratorListRemovedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

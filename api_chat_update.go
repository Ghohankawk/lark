// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// UpdateChat 更新群头像、群名称、群描述、群配置、转让群主等。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 对于群主/群管理员 或 创建群组且具备 [更新应用所创建群的群信息] 权限的机器人, 可更新所有信息
// - 对于不满足上述权限条件的群成员或机器人:
// - 若未开启 [仅群主和群管理员可编辑群信息] 配置, 仅可更新群头像、群名称、群描述、群国际化名称信息
// - 若开启了 [仅群主和群管理员可编辑群信息] 配置, 任何群信息都不能修改
// - 如果同时更新 [邀请用户或机器人入群权限] 和 [群分享权限] 这两项设置需要满足以下条件:
// - 若未开启 [仅群主和管理员可以邀请用户或机器人入群], 需要设置 [群分享权限] 为 [允许分享]
// - 若开启了 [仅群主和管理员可以邀请用户或机器人入群], 需要设置 [群分享权限] 为 [不允许分享]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update
// new doc: https://open.feishu.cn/document/server-docs/group/chat/update-2
func (r *ChatService) UpdateChat(ctx context.Context, request *UpdateChatReq, options ...MethodOptionFunc) (*UpdateChatResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChat mock enable")
		return r.cli.mock.mockChatUpdateChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChat",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChat mock ChatUpdateChat method
func (r *Mock) MockChatUpdateChat(f func(ctx context.Context, request *UpdateChatReq, options ...MethodOptionFunc) (*UpdateChatResp, *Response, error)) {
	r.mockChatUpdateChat = f
}

// UnMockChatUpdateChat un-mock ChatUpdateChat method
func (r *Mock) UnMockChatUpdateChat() {
	r.mockChatUpdateChat = nil
}

// UpdateChatReq ...
type UpdateChatReq struct {
	ChatID                 string                              `path:"chat_id" json:"-"`                   // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`的群组ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	UserIDType             *IDType                             `query:"user_id_type" json:"-"`             // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Avatar                 *string                             `json:"avatar,omitempty"`                   // 群头像对应的 Image Key, 可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意: 上传图片的 [image_type] 需要指定为 [avatar]）, 示例值: "default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string                             `json:"name,omitempty"`                     // 群名称, 示例值: "群聊"
	Description            *string                             `json:"description,omitempty"`              // 群描述, 示例值: "测试群描述"
	I18nNames              *I18nNames                          `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    *AddMemberPermission                `json:"add_member_permission,omitempty"`    // 邀请用户或机器人入群权限, 注意: 若值设置为`only_owner`, 则share_card_permission只能设置为`not_allowed`, 若值设置为`all_members`, 则share_card_permission只能设置为`allowed`, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, 示例值: "all_members"
	ShareCardPermission    *ShareCardPermission                `json:"share_card_permission,omitempty"`    // 群分享权限, 可选值有: `allowed`: 允许, `not_allowed`: 不允许, 示例值: "allowed"
	AtAllPermission        *AtAllPermission                    `json:"at_all_permission,omitempty"`        // at 所有人权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, 示例值: "all_members"
	EditPermission         *EditPermission                     `json:"edit_permission,omitempty"`          // 群编辑权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, 示例值: "all_members"
	OwnerID                *string                             `json:"owner_id,omitempty"`                 // 新群主 ID, 不转让群主时无需填写。ID类型推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), 示例值: "4d7a3c6g"
	JoinMessageVisibility  *MessageVisibility                  `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "only_owner"
	LeaveMessageVisibility *MessageVisibility                  `json:"leave_message_visibility,omitempty"` // 出群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "only_owner"
	MembershipApproval     *MembershipApproval                 `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批, 示例值: "no_approval_required"
	RestrictedModeSetting  *UpdateChatReqRestrictedModeSetting `json:"restricted_mode_setting,omitempty"`  // 保密模式设置
	ChatType               *ChatType                           `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群, 示例值: "private"
}

// UpdateChatReqRestrictedModeSetting ...
type UpdateChatReqRestrictedModeSetting struct {
	Status                         *bool   `json:"status,omitempty"`                            // 保密模式是否开启, 注意: status为true时, screenshot_has_permission_setting、download_has_permission_setting、message_has_permission_setting不能全为all_members, status为false时, screenshot_has_permission_setting、download_has_permission_setting、message_has_permission_setting不能存在not_anyone, 示例值: false
	ScreenshotHasPermissionSetting *string `json:"screenshot_has_permission_setting,omitempty"` // 允许截屏录屏, 示例值: "all_members", 可选值有: all_members: 所有成员允许截屏录屏, not_anyone: 所有成员禁止截屏录屏
	DownloadHasPermissionSetting   *string `json:"download_has_permission_setting,omitempty"`   // 允许下载消息中图片、视频和文件, 示例值: "all_members", 可选值有: all_members: 所有成员允许下载资源, not_anyone: 所有成员禁止下载资源
	MessageHasPermissionSetting    *string `json:"message_has_permission_setting,omitempty"`    // 允许复制和转发消息, 示例值: "all_members", 可选值有: all_members: 所有成员允许复制和转发消息, not_anyone: 所有成员禁止复制和转发消息
}

// UpdateChatResp ...
type UpdateChatResp struct {
}

// updateChatResp ...
type updateChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatResp `json:"data,omitempty"`
}

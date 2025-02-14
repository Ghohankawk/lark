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

// UpdateVCReserveConfigDisableInform 更新禁用状态变更通知
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config-disable_inform/patch
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/scope_config/patch-4
func (r *VCService) UpdateVCReserveConfigDisableInform(ctx context.Context, request *UpdateVCReserveConfigDisableInformReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigDisableInformResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCReserveConfigDisableInform != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCReserveConfigDisableInform mock enable")
		return r.cli.mock.mockVCUpdateVCReserveConfigDisableInform(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "UpdateVCReserveConfigDisableInform",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reserve_configs/:reserve_config_id/disable_inform",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateVCReserveConfigDisableInformResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCReserveConfigDisableInform mock VCUpdateVCReserveConfigDisableInform method
func (r *Mock) MockVCUpdateVCReserveConfigDisableInform(f func(ctx context.Context, request *UpdateVCReserveConfigDisableInformReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigDisableInformResp, *Response, error)) {
	r.mockVCUpdateVCReserveConfigDisableInform = f
}

// UnMockVCUpdateVCReserveConfigDisableInform un-mock VCUpdateVCReserveConfigDisableInform method
func (r *Mock) UnMockVCUpdateVCReserveConfigDisableInform() {
	r.mockVCUpdateVCReserveConfigDisableInform = nil
}

// UpdateVCReserveConfigDisableInformReq ...
type UpdateVCReserveConfigDisableInformReq struct {
	ReserveConfigID string                                              `path:"reserve_config_id" json:"-"` // 会议室或层级ID, 可通过会议室或层级相关查询接口获取, 示例值: "omm_3c5dd7e09bac0c1758fcf9511bd1a771"
	UserIDType      *IDType                                             `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ScopeType       int64                                               `json:"scope_type,omitempty"`       // 1表示会议室层级, 2表示会议室, 示例值: 2, 取值范围: `1` ～ `2`
	DisableInform   *UpdateVCReserveConfigDisableInformReqDisableInform `json:"disable_inform,omitempty"`   // 禁用通知配置
}

// UpdateVCReserveConfigDisableInformReqDisableInform ...
type UpdateVCReserveConfigDisableInformReqDisableInform struct {
	IfInform      bool                                                              `json:"if_inform,omitempty"`      // 禁用状态变更通知开关, 示例值: false, 默认值: `false`
	InformedUsers []*UpdateVCReserveConfigDisableInformReqDisableInformInformedUser `json:"informed_users,omitempty"` // 通知成员列表
	InformedDepts []*UpdateVCReserveConfigDisableInformReqDisableInformInformedDept `json:"informed_depts,omitempty"` // 通知部门列表
}

// UpdateVCReserveConfigDisableInformReqDisableInformInformedDept ...
type UpdateVCReserveConfigDisableInformReqDisableInformInformedDept struct {
	DepartmentID string `json:"department_id,omitempty"` // 预定管理部门ID, 使用open_department_id, 示例值: "od-47d8b570b0a011e9679a755efcc5f61a"
}

// UpdateVCReserveConfigDisableInformReqDisableInformInformedUser ...
type UpdateVCReserveConfigDisableInformReqDisableInformInformedUser struct {
	UserID string `json:"user_id,omitempty"` // 预定管理员ID, 示例值: "ou_a27b07a9071d90577c0177bcec98f856"
}

// UpdateVCReserveConfigDisableInformResp ...
type UpdateVCReserveConfigDisableInformResp struct {
}

// updateVCReserveConfigDisableInformResp ...
type updateVCReserveConfigDisableInformResp struct {
	Code int64                                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCReserveConfigDisableInformResp `json:"data,omitempty"`
}

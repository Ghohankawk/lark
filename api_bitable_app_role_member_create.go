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

// CreateBitableAppRoleMember 新增自定义角色的协作者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/create
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/advanced-permission/app-role-member/create
func (r *BitableService) CreateBitableAppRoleMember(ctx context.Context, request *CreateBitableAppRoleMemberReq, options ...MethodOptionFunc) (*CreateBitableAppRoleMemberResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableAppRoleMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableAppRoleMember mock enable")
		return r.cli.mock.mockBitableCreateBitableAppRoleMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableAppRoleMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableAppRoleMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableAppRoleMember mock BitableCreateBitableAppRoleMember method
func (r *Mock) MockBitableCreateBitableAppRoleMember(f func(ctx context.Context, request *CreateBitableAppRoleMemberReq, options ...MethodOptionFunc) (*CreateBitableAppRoleMemberResp, *Response, error)) {
	r.mockBitableCreateBitableAppRoleMember = f
}

// UnMockBitableCreateBitableAppRoleMember un-mock BitableCreateBitableAppRoleMember method
func (r *Mock) UnMockBitableCreateBitableAppRoleMember() {
	r.mockBitableCreateBitableAppRoleMember = nil
}

// CreateBitableAppRoleMemberReq ...
type CreateBitableAppRoleMemberReq struct {
	AppToken     string  `path:"app_token" json:"-"`       // 多维表格的唯一标识符 [app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	RoleID       string  `path:"role_id" json:"-"`         // 自定义角色的id, 示例值: "roljRpwIUt"
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 协作者id类型, 与请求体中的member_id要对应, 示例值: open_id, 可选值有: open_id: 以open_id来识别协作者, union_id: 以union_id来识别协作者, user_id: 以user_id来识别协作者, chat_id: 以chat_id来识别协作者, department_id: 以department_id来识别协作者, open_department_id: 以open_department_id来识别协作者, 默认值: `open_id`
	MemberID     string  `json:"member_id,omitempty"`      // 协作者id, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

// CreateBitableAppRoleMemberResp ...
type CreateBitableAppRoleMemberResp struct {
}

// createBitableAppRoleMemberResp ...
type createBitableAppRoleMemberResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableAppRoleMemberResp `json:"data,omitempty"`
}

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

// GetAdminBadgeGrantList 通过该接口可以获取特定勋章下的授予名单列表, 授予名单的排列顺序按照创建时间倒序排列。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge-grant/list
func (r *AdminService) GetAdminBadgeGrantList(ctx context.Context, request *GetAdminBadgeGrantListReq, options ...MethodOptionFunc) (*GetAdminBadgeGrantListResp, *Response, error) {
	if r.cli.mock.mockAdminGetAdminBadgeGrantList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#GetAdminBadgeGrantList mock enable")
		return r.cli.mock.mockAdminGetAdminBadgeGrantList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "GetAdminBadgeGrantList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badges/:badge_id/grants",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAdminBadgeGrantListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminGetAdminBadgeGrantList mock AdminGetAdminBadgeGrantList method
func (r *Mock) MockAdminGetAdminBadgeGrantList(f func(ctx context.Context, request *GetAdminBadgeGrantListReq, options ...MethodOptionFunc) (*GetAdminBadgeGrantListResp, *Response, error)) {
	r.mockAdminGetAdminBadgeGrantList = f
}

// UnMockAdminGetAdminBadgeGrantList un-mock AdminGetAdminBadgeGrantList method
func (r *Mock) UnMockAdminGetAdminBadgeGrantList() {
	r.mockAdminGetAdminBadgeGrantList = nil
}

// GetAdminBadgeGrantListReq ...
type GetAdminBadgeGrantListReq struct {
	BadgeID          string            `path:"badge_id" json:"-"`            // 企业勋章的唯一ID, 示例值: "m_DjMzaK", 长度范围: `1` ～ `64` 字符
	PageSize         int64             `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 默认值: `10`, 最大值: `50`
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "om5fn1"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	Name             *string           `query:"name" json:"-"`               // 授予名单名称, 精确匹配, 示例值: "激励勋章的授予名单", 最小长度: `1` 字符
}

// GetAdminBadgeGrantListResp ...
type GetAdminBadgeGrantListResp struct {
	Grants    []*GetAdminBadgeGrantListRespGrant `json:"grants,omitempty"`     // 授予名单列表
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetAdminBadgeGrantListRespGrant ...
type GetAdminBadgeGrantListRespGrant struct {
	ID            string                                     `json:"id,omitempty"`             // 租户内授予名单的唯一标识, 该值由系统随机生成。
	BadgeID       string                                     `json:"badge_id,omitempty"`       // 企业勋章的唯一ID
	Name          string                                     `json:"name,omitempty"`           // 勋章下唯一的授予事项, 最多100个字符。
	GrantType     int64                                      `json:"grant_type,omitempty"`     // 授予名单类型, 可选值有: 0: 手动选择有效期, 1: 匹配系统入职时间
	TimeZone      string                                     `json:"time_zone,omitempty"`      // 授予名单的生效时间对应的时区, 用于检查RuleDetail的时间戳的取值是否规范, 取值范围为TZ database name
	RuleDetail    *GetAdminBadgeGrantListRespGrantRuleDetail `json:"rule_detail,omitempty"`    // 规则详情
	IsGrantAll    bool                                       `json:"is_grant_all,omitempty"`   // 是否授予给全员。1.为false时, 需要关联1~500个用户群体。2.为true时, 不可关联用户、用户组、部门。
	UserIDs       []string                                   `json:"user_ids,omitempty"`       // 授予的用户ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
	DepartmentIDs []string                                   `json:"department_ids,omitempty"` // 授予的部门ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
	GroupIDs      []string                                   `json:"group_ids,omitempty"`      // 授予的用户组ID列表, 授予名单列表接口返回结果中不返回该字段, 只在详情接口返回
}

// GetAdminBadgeGrantListRespGrantRuleDetail ...
type GetAdminBadgeGrantListRespGrantRuleDetail struct {
	EffectiveTime   string `json:"effective_time,omitempty"`   // 开始生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.时间戳必须是所在时区当天的零点时间戳, 如时区为Asia/Shanghai时区时的1649606400
	ExpirationTime  string `json:"expiration_time,omitempty"`  // 结束生效的时间戳。1.手动设置有效期类型勋章, 配置有效期限需要配置该字段；2.最大值: 不得超过effective_time+100 年；3.非永久有效: 时间戳必须是所在时区当天的23:59:59时间戳, 如时区为Asia/Shanghai时区时的1649692799；4.永久有效: 传值为0即可
	Anniversary     int64  `json:"anniversary,omitempty"`      // 入职周年日。根据入职时间发放类型勋章, 需要配置该字段。
	EffectivePeriod int64  `json:"effective_period,omitempty"` // 有效期限。根据入职时间发放类型勋章, 需要配置该字段, 可选值有: 1: 有效期为一年, 2: 永久有效
}

// getAdminBadgeGrantListResp ...
type getAdminBadgeGrantListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetAdminBadgeGrantListResp `json:"data,omitempty"`
}

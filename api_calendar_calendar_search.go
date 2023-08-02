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

// SearchCalendar 该接口用于通过关键字查询公共日历或用户主日历。
//
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/search
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/calendar/search
func (r *CalendarService) SearchCalendar(ctx context.Context, request *SearchCalendarReq, options ...MethodOptionFunc) (*SearchCalendarResp, *Response, error) {
	if r.cli.mock.mockCalendarSearchCalendar != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SearchCalendar mock enable")
		return r.cli.mock.mockCalendarSearchCalendar(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "SearchCalendar",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCalendarResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarSearchCalendar mock CalendarSearchCalendar method
func (r *Mock) MockCalendarSearchCalendar(f func(ctx context.Context, request *SearchCalendarReq, options ...MethodOptionFunc) (*SearchCalendarResp, *Response, error)) {
	r.mockCalendarSearchCalendar = f
}

// UnMockCalendarSearchCalendar un-mock CalendarSearchCalendar method
func (r *Mock) UnMockCalendarSearchCalendar() {
	r.mockCalendarSearchCalendar = nil
}

// SearchCalendarReq ...
type SearchCalendarReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 10
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `50`
	Query     string  `json:"query,omitempty"`      // 搜索关键字, 示例值: "query words", 长度范围: `1` ～ `200` 字符
}

// SearchCalendarResp ...
type SearchCalendarResp struct {
	Items     []*SearchCalendarRespItem `json:"items,omitempty"`      // 搜索命中的日历列表
	PageToken string                    `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// SearchCalendarRespItem ...
type SearchCalendarRespItem struct {
	CalendarID   string             `json:"calendar_id,omitempty"`    // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction)
	Summary      string             `json:"summary,omitempty"`        // 日历标题
	Description  string             `json:"description,omitempty"`    // 日历描述
	Permissions  CalendarPermission `json:"permissions,omitempty"`    // 日历公开范围, 可选值有: private: 私密, show_only_free_busy: 仅展示忙闲信息, public: 他人可查看日程详情
	Color        int64              `json:"color,omitempty"`          // 日历颜色, 颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效
	Type         CalendarType       `json:"type,omitempty"`           // 日历类型, 可选值有: unknown: 未知类型, primary: 用户或应用的主日历, shared: 由用户或应用创建的共享日历, google: 用户绑定的谷歌日历, resource: 会议室日历, exchange: 用户绑定的Exchange日历
	SummaryAlias string             `json:"summary_alias,omitempty"`  // 日历备注名, 修改或添加后仅对当前身份生效
	IsDeleted    bool               `json:"is_deleted,omitempty"`     // 对于当前身份, 日历是否已经被标记为删除
	IsThirdParty bool               `json:"is_third_party,omitempty"` // 当前日历是否是第三方数据；三方日历及日程只支持读, 不支持写入
	Role         CalendarRole       `json:"role,omitempty"`           // 当前身份对于该日历的访问权限, 可选值有: unknown: 未知权限, free_busy_reader: 游客, 只能看到忙碌/空闲信息, reader: 订阅者, 查看所有日程详情, writer: 编辑者, 创建及修改日程, owner: 管理员, 管理日历及共享设置
}

// searchCalendarResp ...
type searchCalendarResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SearchCalendarResp `json:"data,omitempty"`
}

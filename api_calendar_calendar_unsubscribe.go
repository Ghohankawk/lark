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

// UnsubscribeCalendar 该接口用于以当前身份（应用 / 用户）取消对某日历的订阅状态。
//
// 身份由 Header Authorization 的 Token 类型决定。
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 仅可操作已经被当前身份订阅的日历。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscribe
func (r *CalendarService) UnsubscribeCalendar(ctx context.Context, request *UnsubscribeCalendarReq, options ...MethodOptionFunc) (*UnsubscribeCalendarResp, *Response, error) {
	if r.cli.mock.mockCalendarUnsubscribeCalendar != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendar mock enable")
		return r.cli.mock.mockCalendarUnsubscribeCalendar(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "UnsubscribeCalendar",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(unsubscribeCalendarResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarUnsubscribeCalendar mock CalendarUnsubscribeCalendar method
func (r *Mock) MockCalendarUnsubscribeCalendar(f func(ctx context.Context, request *UnsubscribeCalendarReq, options ...MethodOptionFunc) (*UnsubscribeCalendarResp, *Response, error)) {
	r.mockCalendarUnsubscribeCalendar = f
}

// UnMockCalendarUnsubscribeCalendar un-mock CalendarUnsubscribeCalendar method
func (r *Mock) UnMockCalendarUnsubscribeCalendar() {
	r.mockCalendarUnsubscribeCalendar = nil
}

// UnsubscribeCalendarReq ...
type UnsubscribeCalendarReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

// UnsubscribeCalendarResp ...
type UnsubscribeCalendarResp struct {
}

// unsubscribeCalendarResp ...
type unsubscribeCalendarResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UnsubscribeCalendarResp `json:"data,omitempty"`
}

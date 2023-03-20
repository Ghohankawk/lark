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

// DeleteCalendarExchangeBinding 本接口解除Exchange账户和飞书账户的绑定关系, Exchange账户解除绑定后才能绑定其他飞书账户
//
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 操作用户需要是企业超级管理员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/delete
func (r *CalendarService) DeleteCalendarExchangeBinding(ctx context.Context, request *DeleteCalendarExchangeBindingReq, options ...MethodOptionFunc) (*DeleteCalendarExchangeBindingResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarExchangeBinding != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarExchangeBinding mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarExchangeBinding(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "DeleteCalendarExchangeBinding",
		Method:              "DELETE",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(deleteCalendarExchangeBindingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarDeleteCalendarExchangeBinding mock CalendarDeleteCalendarExchangeBinding method
func (r *Mock) MockCalendarDeleteCalendarExchangeBinding(f func(ctx context.Context, request *DeleteCalendarExchangeBindingReq, options ...MethodOptionFunc) (*DeleteCalendarExchangeBindingResp, *Response, error)) {
	r.mockCalendarDeleteCalendarExchangeBinding = f
}

// UnMockCalendarDeleteCalendarExchangeBinding un-mock CalendarDeleteCalendarExchangeBinding method
func (r *Mock) UnMockCalendarDeleteCalendarExchangeBinding() {
	r.mockCalendarDeleteCalendarExchangeBinding = nil
}

// DeleteCalendarExchangeBindingReq ...
type DeleteCalendarExchangeBindingReq struct {
	ExchangeBindingID string `path:"exchange_binding_id" json:"-"` // exchange绑定唯一标识id。参见[exchange绑定ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/introduction#12533d5e), 示例值: "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20="
}

// DeleteCalendarExchangeBindingResp ...
type DeleteCalendarExchangeBindingResp struct {
}

// deleteCalendarExchangeBindingResp ...
type deleteCalendarExchangeBindingResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarExchangeBindingResp `json:"data,omitempty"`
}

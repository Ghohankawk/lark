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

// BatchDeleteMessage 批量撤回通过[批量发送消息](https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口发送的消息。
//
// 注意事项:
// - 应用需要启用[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 撤回单条发送的消息请使用[撤回消息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete)接口
// - 不支持撤回发出时间超过1天的消息
// - 一次调用涉及大量消息, 所以为异步接口, 会有一定延迟
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/delete
// new doc: https://open.feishu.cn/document/server-docs/im-v1/batch_message/delete
func (r *MessageService) BatchDeleteMessage(ctx context.Context, request *BatchDeleteMessageReq, options ...MethodOptionFunc) (*BatchDeleteMessageResp, *Response, error) {
	if r.cli.mock.mockMessageBatchDeleteMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#BatchDeleteMessage mock enable")
		return r.cli.mock.mockMessageBatchDeleteMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "BatchDeleteMessage",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/batch_messages/:batch_message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeleteMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageBatchDeleteMessage mock MessageBatchDeleteMessage method
func (r *Mock) MockMessageBatchDeleteMessage(f func(ctx context.Context, request *BatchDeleteMessageReq, options ...MethodOptionFunc) (*BatchDeleteMessageResp, *Response, error)) {
	r.mockMessageBatchDeleteMessage = f
}

// UnMockMessageBatchDeleteMessage un-mock MessageBatchDeleteMessage method
func (r *Mock) UnMockMessageBatchDeleteMessage() {
	r.mockMessageBatchDeleteMessage = nil
}

// BatchDeleteMessageReq ...
type BatchDeleteMessageReq struct {
	BatchMessageID string `path:"batch_message_id" json:"-"` // 待撤回的批量消息的ID, 为[批量发送消息](https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口返回值中的`message_id`字段, 用于标识一次批量发送消息请求, 示例值: "bm-dc13264520392913993dd051dba21dcf"
}

// BatchDeleteMessageResp ...
type BatchDeleteMessageResp struct {
}

// batchDeleteMessageResp ...
type batchDeleteMessageResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteMessageResp `json:"data,omitempty"`
}

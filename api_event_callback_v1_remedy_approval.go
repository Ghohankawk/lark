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

// EventV1RemedyApproval 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 补卡申请审批通过后触发此事件。 你可以在「打卡」应用里提交补卡申请。
// * 依赖权限: [访问审批应用] 或 [查看、创建、更新、删除审批应用相关信息]
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/attendance-record-correction
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/event/special-event/attendance-record-correction
func (r *EventCallbackService) HandlerEventV1RemedyApproval(f EventV1RemedyApprovalHandler) {
	r.cli.eventHandler.eventV1RemedyApprovalHandler = f
}

// EventV1RemedyApprovalHandler event EventV1RemedyApproval handler
type EventV1RemedyApprovalHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1RemedyApproval) (string, error)

// EventV1RemedyApproval ...
type EventV1RemedyApproval struct {
	Object *EventV1RemedyApprovalObject `json:"object,omitempty"` // 为当前的数据, 事件的标准格式
}

// EventV1RemedyApprovalObject ...
type EventV1RemedyApprovalObject struct {
	Type         string `json:"type,omitempty"`          // 类型. 如: remedy_approval_v2
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例code
	EmployeeID   string `json:"employee_id,omitempty"`   // 用户id
	StartTime    int64  `json:"start_time,omitempty"`    // 审批发起时间. 如: 0
	EndTime      int64  `json:"end_time,omitempty"`      // 审批结束时间. 如: 0
	RemedyTime   int64  `json:"remedy_time,omitempty"`   // 补卡时间. 如: 0
	RemedyReason string `json:"remedy_reason,omitempty"` // 补卡原因
	Status       string `json:"status,omitempty"`        // 实例状态
}

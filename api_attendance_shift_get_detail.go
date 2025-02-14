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

// GetAttendanceShiftDetail 通过班次 ID 获取班次详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/get
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/shift/get
func (r *AttendanceService) GetAttendanceShiftDetail(ctx context.Context, request *GetAttendanceShiftDetailReq, options ...MethodOptionFunc) (*GetAttendanceShiftDetailResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceShiftDetail != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceShiftDetail mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceShiftDetail(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceShiftDetail",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceShiftDetailResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceShiftDetail mock AttendanceGetAttendanceShiftDetail method
func (r *Mock) MockAttendanceGetAttendanceShiftDetail(f func(ctx context.Context, request *GetAttendanceShiftDetailReq, options ...MethodOptionFunc) (*GetAttendanceShiftDetailResp, *Response, error)) {
	r.mockAttendanceGetAttendanceShiftDetail = f
}

// UnMockAttendanceGetAttendanceShiftDetail un-mock AttendanceGetAttendanceShiftDetail method
func (r *Mock) UnMockAttendanceGetAttendanceShiftDetail() {
	r.mockAttendanceGetAttendanceShiftDetail = nil
}

// GetAttendanceShiftDetailReq ...
type GetAttendanceShiftDetailReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID, 获取方式: 1）[按名称查询班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query) 2）[创建班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create), 示例值: "6919358778597097404"
}

// GetAttendanceShiftDetailResp ...
type GetAttendanceShiftDetailResp struct {
	ShiftID           string                                           `json:"shift_id,omitempty"`              // 班次 ID
	ShiftName         string                                           `json:"shift_name,omitempty"`            // 班次名称
	PunchTimes        int64                                            `json:"punch_times,omitempty"`           // 打卡次数
	IsFlexible        bool                                             `json:"is_flexible,omitempty"`           // 是否弹性打卡
	FlexibleMinutes   int64                                            `json:"flexible_minutes,omitempty"`      // 弹性打卡时间, 设置【上班最多可晚到】与【下班最多可早走】时间, 如果不设置flexible_rule则生效
	FlexibleRule      []*GetAttendanceShiftDetailRespFlexibleRule      `json:"flexible_rule,omitempty"`         // 弹性打卡时间设置
	NoNeedOff         bool                                             `json:"no_need_off,omitempty"`           // 不需要打下班卡
	PunchTimeRule     []*GetAttendanceShiftDetailRespPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule []*GetAttendanceShiftDetailRespLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      []*GetAttendanceShiftDetailRespRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息规则
	OvertimeRule      []*GetAttendanceShiftDetailRespOvertimeRule      `json:"overtime_rule,omitempty"`         // 打卡规则
}

// GetAttendanceShiftDetailRespFlexibleRule ...
type GetAttendanceShiftDetailRespFlexibleRule struct {
	FlexibleEarlyMinutes int64 `json:"flexible_early_minutes,omitempty"` // 下班最多可早走（上班早到几分钟, 下班可早走几分钟）
	FlexibleLateMinutes  int64 `json:"flexible_late_minutes,omitempty"`  // 上班最多可晚到（上班晚到几分钟, 下班须晚走几分钟）
}

// GetAttendanceShiftDetailRespLateOffLateOnRule ...
type GetAttendanceShiftDetailRespLateOffLateOnRule struct {
	LateOffMinutes int64 `json:"late_off_minutes,omitempty"` // 晚走多久
	LateOnMinutes  int64 `json:"late_on_minutes,omitempty"`  // 晚到多久
}

// GetAttendanceShiftDetailRespOvertimeRule ...
type GetAttendanceShiftDetailRespOvertimeRule struct {
	OnOvertime  string `json:"on_overtime,omitempty"`  // 上班时间
	OffOvertime string `json:"off_overtime,omitempty"` // 下班时间
}

// GetAttendanceShiftDetailRespPunchTimeRule ...
type GetAttendanceShiftDetailRespPunchTimeRule struct {
	OnTime                   string `json:"on_time,omitempty"`                      // 上班时间
	OffTime                  string `json:"off_time,omitempty"`                     // 下班时间
	LateMinutesAsLate        int64  `json:"late_minutes_as_late,omitempty"`         // 晚到多久记为迟到
	LateMinutesAsLack        int64  `json:"late_minutes_as_lack,omitempty"`         // 晚到多久记为缺卡
	OnAdvanceMinutes         int64  `json:"on_advance_minutes,omitempty"`           // 最早多久可打上班卡
	EarlyMinutesAsEarly      int64  `json:"early_minutes_as_early,omitempty"`       // 早退多久记为早退
	EarlyMinutesAsLack       int64  `json:"early_minutes_as_lack,omitempty"`        // 早退多久记为缺卡
	OffDelayMinutes          int64  `json:"off_delay_minutes,omitempty"`            // 最晚多久可打下班卡
	LateMinutesAsSeriousLate int64  `json:"late_minutes_as_serious_late,omitempty"` // 晚到多久记为严重迟到
}

// GetAttendanceShiftDetailRespRestTimeRule ...
type GetAttendanceShiftDetailRespRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束
}

// getAttendanceShiftDetailResp ...
type getAttendanceShiftDetailResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceShiftDetailResp `json:"data,omitempty"`
}

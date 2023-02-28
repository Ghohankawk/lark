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

// GetVCParticipantQualityList 查询参会人会议质量数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/participant_quality_list/get
func (r *VCService) GetVCParticipantQualityList(ctx context.Context, request *GetVCParticipantQualityListReq, options ...MethodOptionFunc) (*GetVCParticipantQualityListResp, *Response, error) {
	if r.cli.mock.mockVCGetVCParticipantQualityList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCParticipantQualityList mock enable")
		return r.cli.mock.mockVCGetVCParticipantQualityList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCParticipantQualityList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/participant_quality_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCParticipantQualityListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCParticipantQualityList mock VCGetVCParticipantQualityList method
func (r *Mock) MockVCGetVCParticipantQualityList(f func(ctx context.Context, request *GetVCParticipantQualityListReq, options ...MethodOptionFunc) (*GetVCParticipantQualityListResp, *Response, error)) {
	r.mockVCGetVCParticipantQualityList = f
}

// UnMockVCGetVCParticipantQualityList un-mock VCGetVCParticipantQualityList method
func (r *Mock) UnMockVCGetVCParticipantQualityList() {
	r.mockVCGetVCParticipantQualityList = nil
}

// GetVCParticipantQualityListReq ...
type GetVCParticipantQualityListReq struct {
	MeetingStartTime string  `query:"meeting_start_time" json:"-"` // 会议开始时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingEndTime   string  `query:"meeting_end_time" json:"-"`   // 会议结束时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingNo        string  `query:"meeting_no" json:"-"`         // 9位会议号, 示例值: "123456789"
	JoinTime         string  `query:"join_time" json:"-"`          // 参会人入会时间（unix时间, 单位sec）, 示例值: "1655276858"
	UserID           *string `query:"user_id" json:"-"`            // 参会人为Lark用户时填入, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	RoomID           *string `query:"room_id" json:"-"`            // 参会人为Rooms时填入, 示例值: "omm_eada1d61a550955240c28757e7dec3af"
	PageSize         *int64  `query:"page_size" json:"-"`          // 分页大小, 示例值: 20, 默认值: `20`, 最大值: `100`
	PageToken        *string `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "20"
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCParticipantQualityListResp ...
type GetVCParticipantQualityListResp struct {
	ParticipantQualityList []*GetVCParticipantQualityListRespParticipantQuality `json:"participant_quality_list,omitempty"` // 参会人参会质量列表
	PageToken              string                                               `json:"page_token,omitempty"`               // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore                bool                                                 `json:"has_more,omitempty"`                 // 是否还有更多项
}

// GetVCParticipantQualityListRespParticipantQuality ...
type GetVCParticipantQualityListRespParticipantQuality struct {
	Network       *GetVCParticipantQualityListRespParticipantQualityNetwork       `json:"network,omitempty"`        // 网络
	Audio         *GetVCParticipantQualityListRespParticipantQualityAudio         `json:"audio,omitempty"`          // 音频
	Video         *GetVCParticipantQualityListRespParticipantQualityVideo         `json:"video,omitempty"`          // 视频
	ScreenSharing *GetVCParticipantQualityListRespParticipantQualityScreenSharing `json:"screen_sharing,omitempty"` // 共享屏幕
	CpuUsage      *GetVCParticipantQualityListRespParticipantQualityCpuUsage      `json:"cpu_usage,omitempty"`      // Cpu使用量
}

// GetVCParticipantQualityListRespParticipantQualityAudio ...
type GetVCParticipantQualityListRespParticipantQualityAudio struct {
	Time            string `json:"time,omitempty"`             // 时间
	MicInputVolume  string `json:"mic_input_volume,omitempty"` // 麦克风采集音量
	SpeakerVolume   string `json:"speaker_volume,omitempty"`   // 扬声器播放音量
	BitrateReceived string `json:"bitrate_received,omitempty"` // 码率（接收）
	LatencyReceived string `json:"latency_received,omitempty"` // 延迟（接收）
	JitterReceived  string `json:"jitter_received,omitempty"`  // 抖动（接收）
	BitrateSent     string `json:"bitrate_sent,omitempty"`     // 码率（发送）
	LatencySent     string `json:"latency_sent,omitempty"`     // 延迟（发送）
	JitterSent      string `json:"jitter_sent,omitempty"`      // 抖动（发送）
}

// GetVCParticipantQualityListRespParticipantQualityCpuUsage ...
type GetVCParticipantQualityListRespParticipantQualityCpuUsage struct {
	Time              string `json:"time,omitempty"`                 // 时间
	ClientAvgCpuUsage string `json:"client_avg_cpu_usage,omitempty"` // 客户端平均 CPU 占用
	ClientMaxCpuUsage string `json:"client_max_cpu_usage,omitempty"` // 客户端最大 CPU 占用
	SystemAvgCpuUsage string `json:"system_avg_cpu_usage,omitempty"` // 系统平均 CPU 占用
	SystemMaxCpuUsage string `json:"system_max_cpu_usage,omitempty"` // 系统最大 CPU 占用
}

// GetVCParticipantQualityListRespParticipantQualityNetwork ...
type GetVCParticipantQualityListRespParticipantQualityNetwork struct {
	Time                  string `json:"time,omitempty"`                     // 时间
	NetworkDelay          string `json:"network_delay,omitempty"`            // 网络延迟
	BitrateReceived       string `json:"bitrate_received,omitempty"`         // 码率（接收）
	PacketLossAvgReceived string `json:"packet_loss_avg_received,omitempty"` // 丢包 - 平均（接收）
	PacketLossMaxReceived string `json:"packet_loss_max_received,omitempty"` // 丢包 - 最大（接收）
	BitrateSent           string `json:"bitrate_sent,omitempty"`             // 码率（发送）
	PacketLossAvgSent     string `json:"packet_loss_avg_sent,omitempty"`     // 丢包 - 平均（发送）
	PacketLossMaxSent     string `json:"packet_loss_max_sent,omitempty"`     // 丢包 - 最大（发送）
}

// GetVCParticipantQualityListRespParticipantQualityScreenSharing ...
type GetVCParticipantQualityListRespParticipantQualityScreenSharing struct {
	Time                      string `json:"time,omitempty"`                        // 时间
	BitrateReceived           string `json:"bitrate_received,omitempty"`            // 码率（接收）
	LatencyReceived           string `json:"latency_received,omitempty"`            // 延迟（接收）
	JitterReceived            string `json:"jitter_received,omitempty"`             // 抖动（接收）
	MaximumResolutionReceived string `json:"maximum_resolution_received,omitempty"` // 最大分辨率（接收）
	FramerateReceived         string `json:"framerate_received,omitempty"`          // 帧率（接收）
	BitrateSent               string `json:"bitrate_sent,omitempty"`                // 码率（发送）
	LatencySent               string `json:"latency_sent,omitempty"`                // 延迟（发送）
	JitterSent                string `json:"jitter_sent,omitempty"`                 // 抖动（发送）
	MaximumResolutionSent     string `json:"maximum_resolution_sent,omitempty"`     // 最大分辨率（发送）
	FramerateSent             string `json:"framerate_sent,omitempty"`              // 帧率（发送）
}

// GetVCParticipantQualityListRespParticipantQualityVideo ...
type GetVCParticipantQualityListRespParticipantQualityVideo struct {
	Time                      string `json:"time,omitempty"`                        // 时间
	BitrateReceived           string `json:"bitrate_received,omitempty"`            // 码率（接收）
	LatencyReceived           string `json:"latency_received,omitempty"`            // 延迟（接收）
	JitterReceived            string `json:"jitter_received,omitempty"`             // 抖动（接收）
	MaximumResolutionReceived string `json:"maximum_resolution_received,omitempty"` // 最大分辨率（接收）
	FramerateReceived         string `json:"framerate_received,omitempty"`          // 帧率（接收）
	BitrateSent               string `json:"bitrate_sent,omitempty"`                // 码率（发送）
	LatencySent               string `json:"latency_sent,omitempty"`                // 延迟（发送）
	JitterSent                string `json:"jitter_sent,omitempty"`                 // 抖动（发送）
	MaximumResolutionSent     string `json:"maximum_resolution_sent,omitempty"`     // 最大分辨率（发送）
	FramerateSent             string `json:"framerate_sent,omitempty"`              // 帧率（发送）
}

// getVCParticipantQualityListResp ...
type getVCParticipantQualityListResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetVCParticipantQualityListResp `json:"data,omitempty"`
}
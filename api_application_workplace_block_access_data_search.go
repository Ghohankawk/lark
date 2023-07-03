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

// SearchApplicationWorkplaceBlockAccessData 获取定制工作台小组件(block)访问数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/workplace-v1/workplace_block_access_data/search
func (r *ApplicationService) SearchApplicationWorkplaceBlockAccessData(ctx context.Context, request *SearchApplicationWorkplaceBlockAccessDataReq, options ...MethodOptionFunc) (*SearchApplicationWorkplaceBlockAccessDataResp, *Response, error) {
	if r.cli.mock.mockApplicationSearchApplicationWorkplaceBlockAccessData != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#SearchApplicationWorkplaceBlockAccessData mock enable")
		return r.cli.mock.mockApplicationSearchApplicationWorkplaceBlockAccessData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "SearchApplicationWorkplaceBlockAccessData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/workplace/v1/workplace_block_access_data/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchApplicationWorkplaceBlockAccessDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationSearchApplicationWorkplaceBlockAccessData mock ApplicationSearchApplicationWorkplaceBlockAccessData method
func (r *Mock) MockApplicationSearchApplicationWorkplaceBlockAccessData(f func(ctx context.Context, request *SearchApplicationWorkplaceBlockAccessDataReq, options ...MethodOptionFunc) (*SearchApplicationWorkplaceBlockAccessDataResp, *Response, error)) {
	r.mockApplicationSearchApplicationWorkplaceBlockAccessData = f
}

// UnMockApplicationSearchApplicationWorkplaceBlockAccessData un-mock ApplicationSearchApplicationWorkplaceBlockAccessData method
func (r *Mock) UnMockApplicationSearchApplicationWorkplaceBlockAccessData() {
	r.mockApplicationSearchApplicationWorkplaceBlockAccessData = nil
}

// SearchApplicationWorkplaceBlockAccessDataReq ...
type SearchApplicationWorkplaceBlockAccessDataReq struct {
	FromDate  string  `query:"from_date" json:"-"`  // 数据检索开始时间, 精确到日。格式yyyy-MM-dd, 示例值: 2023-02-01
	ToDate    string  `query:"to_date" json:"-"`    // 数据检索结束时间, 精确到日。格式yyyy-MM-dd, 示例值: 2023-03-02
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 最小为 1, 最大为 200, 默认为 20, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `200`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: ddowkdkl9w2d
	BlockID   *string `query:"block_id" json:"-"`   // 小组件id（BlockID）, 可前往 飞书管理后台 > 工作台 > 定制工作台, 选择指定的工作台并进入工作台编辑器, 点击某个小组件, 即可查看页面右侧面板中该小组件名称下方的“BlockID”, 示例值: 283438293839422334
}

// SearchApplicationWorkplaceBlockAccessDataResp ...
type SearchApplicationWorkplaceBlockAccessDataResp struct {
	Items     []*DocxBlock `json:"items,omitempty"`      // 工作台中block的访问数据
	HasMore   bool         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// searchApplicationWorkplaceBlockAccessDataResp ...
type searchApplicationWorkplaceBlockAccessDataResp struct {
	Code int64                                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                         `json:"msg,omitempty"`  // 错误描述
	Data *SearchApplicationWorkplaceBlockAccessDataResp `json:"data,omitempty"`
}

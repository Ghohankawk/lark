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

// UpdateSearchDataSource 更新一个已经存在的数据源。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/patch
func (r *SearchService) UpdateSearchDataSource(ctx context.Context, request *UpdateSearchDataSourceReq, options ...MethodOptionFunc) (*UpdateSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchUpdateSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#UpdateSearchDataSource mock enable")
		return r.cli.mock.mockSearchUpdateSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "UpdateSearchDataSource",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchUpdateSearchDataSource mock SearchUpdateSearchDataSource method
func (r *Mock) MockSearchUpdateSearchDataSource(f func(ctx context.Context, request *UpdateSearchDataSourceReq, options ...MethodOptionFunc) (*UpdateSearchDataSourceResp, *Response, error)) {
	r.mockSearchUpdateSearchDataSource = f
}

// UnMockSearchUpdateSearchDataSource un-mock SearchUpdateSearchDataSource method
func (r *Mock) UnMockSearchUpdateSearchDataSource() {
	r.mockSearchUpdateSearchDataSource = nil
}

// UpdateSearchDataSourceReq ...
type UpdateSearchDataSourceReq struct {
	DataSourceID    string                                    `path:"data_source_id" json:"-"`    // 数据源的唯一标识, 示例值: "service_ticket"
	Name            *string                                   `json:"name,omitempty"`             // 数据源的展示名称, 示例值: "客服工单"
	State           *int64                                    `json:"state,omitempty"`            // 数据源状态, 0-已上线, 1-未上线, 示例值: 0, 可选值有: 0: 已上线, 1: 未上线
	Description     *string                                   `json:"description,omitempty"`      // 对于数据源的描述, 示例值: "搜索客服工单"
	IconURL         *string                                   `json:"icon_url,omitempty"`         // 数据源在 search tab 上的展示图标路径, 示例值: "https://www.xxx.com/open.jpg"
	I18nName        *UpdateSearchDataSourceReqI18nName        `json:"i18n_name,omitempty"`        // 数据源名称多语言配置, json格式, key为语言locale, value为对应文案, 例如{"zh_cn":"测试数据源", "en_us":"Test DataSource"}
	I18nDescription *UpdateSearchDataSourceReqI18nDescription `json:"i18n_description,omitempty"` // 数据源描述多语言配置, json格式, key为语言locale, value为对应文案, 例如{"zh_cn":"搜索测试数据源相关数据", "en_us":"Search data from Test DataSource"}
}

// UpdateSearchDataSourceReqI18nDescription ...
type UpdateSearchDataSourceReqI18nDescription struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 国际化字段: 中文, 示例值: "任务"
	EnUs *string `json:"en_us,omitempty"` // 国际化字段: 英文, 示例值: "TODO"
	JaJp *string `json:"ja_jp,omitempty"` // 国际化字段: 日文, 示例值: "タスク"
}

// UpdateSearchDataSourceReqI18nName ...
type UpdateSearchDataSourceReqI18nName struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 国际化字段: 中文, 示例值: "任务"
	EnUs *string `json:"en_us,omitempty"` // 国际化字段: 英文, 示例值: "TODO"
	JaJp *string `json:"ja_jp,omitempty"` // 国际化字段: 日文, 示例值: "タスク"
}

// UpdateSearchDataSourceResp ...
type UpdateSearchDataSourceResp struct {
	DataSource *UpdateSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源
}

// UpdateSearchDataSourceRespDataSource ...
type UpdateSearchDataSourceRespDataSource struct {
	ID               string                                               `json:"id,omitempty"`                // 数据源的唯一标识
	Name             string                                               `json:"name,omitempty"`              // data_source的展示名称
	State            int64                                                `json:"state,omitempty"`             // 数据源状态, 0-已上线, 1-未上线。如果未填, 默认是未上线状态, 可选值有: 0: 已上线, 1: 未上线
	Description      string                                               `json:"description,omitempty"`       // 对于数据源的描述
	CreateTime       string                                               `json:"create_time,omitempty"`       // 创建时间, 使用Unix时间戳, 单位为“秒”
	UpdateTime       string                                               `json:"update_time,omitempty"`       // 更新时间, 使用Unix时间戳, 单位为“秒”
	IsExceedQuota    bool                                                 `json:"is_exceed_quota,omitempty"`   // 是否超限
	IconURL          string                                               `json:"icon_url,omitempty"`          // 数据源在 search tab 上的展示图标路径, 建议使用png或jpeg格式, 否则可能无法在客户端正常展示
	Template         string                                               `json:"template,omitempty"`          // 数据源采用的展示模版名称
	SearchableFields []string                                             `json:"searchable_fields,omitempty"` // 【已废弃, 如有定制需要请使用“数据范式”接口】描述哪些字段可以被搜索
	I18nName         *UpdateSearchDataSourceRespDataSourceI18nName        `json:"i18n_name,omitempty"`         // 数据源的国际化展示名称
	I18nDescription  *UpdateSearchDataSourceRespDataSourceI18nDescription `json:"i18n_description,omitempty"`  // 数据源的国际化描述
	SchemaID         string                                               `json:"schema_id,omitempty"`         // 数据源关联的 schema 标识
}

// UpdateSearchDataSourceRespDataSourceI18nDescription ...
type UpdateSearchDataSourceRespDataSourceI18nDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// UpdateSearchDataSourceRespDataSourceI18nName ...
type UpdateSearchDataSourceRespDataSourceI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// updateSearchDataSourceResp ...
type updateSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSearchDataSourceResp `json:"data,omitempty"`
}

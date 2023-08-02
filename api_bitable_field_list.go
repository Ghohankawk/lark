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

// GetBitableFieldList 根据 app_token 和 table_id, 获取数据表的所有字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-field/list
func (r *BitableService) GetBitableFieldList(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableFieldList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableFieldList mock enable")
		return r.cli.mock.mockBitableGetBitableFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableFieldList mock BitableGetBitableFieldList method
func (r *Mock) MockBitableGetBitableFieldList(f func(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error)) {
	r.mockBitableGetBitableFieldList = f
}

// UnMockBitableGetBitableFieldList un-mock BitableGetBitableFieldList method
func (r *Mock) UnMockBitableGetBitableFieldList() {
	r.mockBitableGetBitableFieldList = nil
}

// GetBitableFieldListReq ...
type GetBitableFieldListReq struct {
	AppToken         string  `path:"app_token" json:"-"`            // Base app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID          string  `path:"table_id" json:"-"`             // table id, 示例值: "tblsRc9GRRXKqhvW"
	ViewID           *string `query:"view_id" json:"-"`             // 视图 ID, 示例值: vewOVMEXPF
	TextFieldAsArray *bool   `query:"text_field_as_array" json:"-"` // 控制字段描述（多行文本格式）数据的返回格式, true 表示以数组富文本形式返回, 示例值: true
	PageToken        *string `query:"page_token" json:"-"`          // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: fldwJ4YrtB
	PageSize         *int64  `query:"page_size" json:"-"`           // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `100`
}

// GetBitableFieldListResp ...
type GetBitableFieldListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Total     int64                          `json:"total,omitempty"`      // 总数
	Items     []*GetBitableFieldListRespItem `json:"items,omitempty"`      // 字段信息
}

// GetBitableFieldListRespItem ...
type GetBitableFieldListRespItem struct {
	FieldID     string                                  `json:"field_id,omitempty"`    // 多维表格字段 id
	FieldName   string                                  `json:"field_name,omitempty"`  // 多维表格字段名, 请注意: 1. 名称中的首尾空格将会被去除。
	Type        int64                                   `json:"type,omitempty"`        // 多维表格字段类型, 可选值有: 1: 多行文本, 2: 数字, 3: 单选, 4: 多选, 5: 日期, 7: 复选框, 11: 人员, 15: 超链接, 17: 附件, 18: 关联, 20: 公式, 21: 双向关联, 1001: 创建时间, 1002: 最后更新时间, 1003: 创建人, 1004: 修改人, 1005: 自动编号, 13: 电话号码, 22: 地理位置, 23: 群组
	Property    *GetBitableFieldListRespItemProperty    `json:"property,omitempty"`    // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
	Description *GetBitableFieldListRespItemDescription `json:"description,omitempty"` // 字段的描述
	IsPrimary   bool                                    `json:"is_primary,omitempty"`  // 是否是索引列
	UiType      string                                  `json:"ui_type,omitempty"`     // 字段在界面上的展示类型, 例如进度字段是数字的一种展示形态
	IsHidden    bool                                    `json:"is_hidden,omitempty"`   // 是否是隐藏字段
}

// GetBitableFieldListRespItemDescription ...
type GetBitableFieldListRespItemDescription struct {
	DisableSync bool   `json:"disable_sync,omitempty"` // 是否禁止同步, 如果为true, 表示禁止同步该描述内容到表单的问题描述（只在新增、修改字段时生效）
	Text        string `json:"text,omitempty"`         // 字段描述内容
}

// GetBitableFieldListRespItemProperty ...
type GetBitableFieldListRespItemProperty struct {
	Options           []*GetBitableFieldListRespItemPropertyOption         `json:"options,omitempty"`            // 单选、多选字段的选项信息
	Formatter         string                                               `json:"formatter,omitempty"`          // 数字、公式字段的显示格式
	DateFormatter     string                                               `json:"date_formatter,omitempty"`     // 日期、创建时间、最后更新时间字段的显示格式
	AutoFill          bool                                                 `json:"auto_fill,omitempty"`          // 日期字段中新纪录自动填写创建时间
	Multiple          bool                                                 `json:"multiple,omitempty"`           // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录
	TableID           string                                               `json:"table_id,omitempty"`           // 单向关联、双向关联字段中关联的数据表的id
	TableName         string                                               `json:"table_name,omitempty"`         // 单向关联、双向关联字段中关联的数据表的名字
	BackFieldName     string                                               `json:"back_field_name,omitempty"`    // 双向关联字段中关联的数据表中对应的双向关联字段的名字
	AutoSerial        *GetBitableFieldListRespItemPropertyAutoSerial       `json:"auto_serial,omitempty"`        // 自动编号类型
	Location          *GetBitableFieldListRespItemPropertyLocation         `json:"location,omitempty"`           // 地理位置输入方式
	FormulaExpression string                                               `json:"formula_expression,omitempty"` // 公式字段的表达式
	AllowedEditModes  *GetBitableFieldListRespItemPropertyAllowedEditModes `json:"allowed_edit_modes,omitempty"` // 字段支持的编辑模式
	Min               float64                                              `json:"min,omitempty"`                // 进度、评分等字段的数据范围最小值
	Max               float64                                              `json:"max,omitempty"`                // 进度、评分等字段的数据范围最大值
	RangeCustomize    bool                                                 `json:"range_customize,omitempty"`    // 进度等字段是否支持自定义范围
	CurrencyCode      string                                               `json:"currency_code,omitempty"`      // 货币币种
	Rating            *GetBitableFieldListRespItemPropertyRating           `json:"rating,omitempty"`             // 评分字段的相关设置
}

// GetBitableFieldListRespItemPropertyAllowedEditModes ...
type GetBitableFieldListRespItemPropertyAllowedEditModes struct {
	Manual bool `json:"manual,omitempty"` // 是否允许手动录入
	Scan   bool `json:"scan,omitempty"`   // 是否允许移动端录入
}

// GetBitableFieldListRespItemPropertyAutoSerial ...
type GetBitableFieldListRespItemPropertyAutoSerial struct {
	Type    string                                                 `json:"type,omitempty"`    // 自动编号类型, 可选值有: custom: 自定义编号, auto_increment_number: 自增数字
	Options []*GetBitableFieldListRespItemPropertyAutoSerialOption `json:"options,omitempty"` // 自动编号规则列表
}

// GetBitableFieldListRespItemPropertyAutoSerialOption ...
type GetBitableFieldListRespItemPropertyAutoSerialOption struct {
	Type  string `json:"type,omitempty"`  // 自动编号的可选规则项类型, 可选值有: system_number: 自增数字位, value范围1-9, fixed_text: 固定字符, 最大长度: 20, created_time: 创建时间, 支持格式 "yyyyMMdd"、"yyyyMM"、"yyyy"、"MMdd"、"MM"、"dd"
	Value string `json:"value,omitempty"` // 与自动编号的可选规则项类型相对应的取值
}

// GetBitableFieldListRespItemPropertyLocation ...
type GetBitableFieldListRespItemPropertyLocation struct {
	InputType string `json:"input_type,omitempty"` // 地理位置输入限制, 可选值有: only_mobile: 只允许移动端上传, not_limit: 无限制
}

// GetBitableFieldListRespItemPropertyOption ...
type GetBitableFieldListRespItemPropertyOption struct {
	Name  string `json:"name,omitempty"`  // 选项名
	ID    string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID
	Color int64  `json:"color,omitempty"` // 选项颜色
}

// GetBitableFieldListRespItemPropertyRating ...
type GetBitableFieldListRespItemPropertyRating struct {
	Symbol string `json:"symbol,omitempty"` // 评分字段的符号展示
}

// getBitableFieldListResp ...
type getBitableFieldListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableFieldListResp `json:"data,omitempty"`
}

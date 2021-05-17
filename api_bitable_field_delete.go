// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteField 该接口用于在数据表中删除一个字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete
func (r *BitableService) DeleteField(ctx context.Context, request *DeleteFieldReq, options ...MethodOptionFunc) (*DeleteFieldResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteField != nil {
		r.cli.logDebug(ctx, "[lark] Bitable#DeleteField mock enable")
		return r.cli.mock.mockBitableDeleteField(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Bitable#DeleteField call api")
	r.cli.logDebug(ctx, "[lark] Bitable#DeleteField request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "DELETE",
		URL:                 "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(deleteFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Bitable#DeleteField DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Bitable#DeleteField DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "DeleteField", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Bitable#DeleteField request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableDeleteField(f func(ctx context.Context, request *DeleteFieldReq, options ...MethodOptionFunc) (*DeleteFieldResp, *Response, error)) {
	r.mockBitableDeleteField = f
}

func (r *Mock) UnMockBitableDeleteField() {
	r.mockBitableDeleteField = nil
}

type DeleteFieldReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值："tblsRc9GRRXKqhvW"
	FieldID  string `path:"field_id" json:"-"`  // field id, 示例值："fldPTb0U2y"
}

type deleteFieldResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteFieldResp `json:"data,omitempty"` //
}

type DeleteFieldResp struct {
	FieldID string `json:"field_id,omitempty"` // field id
	Deleted bool   `json:"deleted,omitempty"`  // 删除标记
}
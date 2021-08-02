// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"github.com/chyroc/lark/internal"
)

// OpenLark 唤起飞书客户端
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-lark
func (r *AppLinkService) OpenLark(req *OpenLarkReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/op/open", req)
}

type OpenLarkReq struct{}
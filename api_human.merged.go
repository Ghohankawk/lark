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
	"io"
)

// GetFaceVerifyAuthResult
//
// 无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
// 无源人脸比对流程，开发者后台通过调用此接口请求飞书后台，对本次活体比对结果做校验。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/query-recognition-result
func (r *HumanAuthService) GetFaceVerifyAuthResult(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error) {
	if r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#GetFaceVerifyAuthResult mock enable")
		return r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "GetFaceVerifyAuthResult",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/face_verify/v1/query_auth_result",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getFaceVerifyAuthResultResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthGetFaceVerifyAuthResult mock HumanAuthGetFaceVerifyAuthResult method
func (r *Mock) MockHumanAuthGetFaceVerifyAuthResult(f func(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error)) {
	r.mockHumanAuthGetFaceVerifyAuthResult = f
}

// UnMockHumanAuthGetFaceVerifyAuthResult un-mock HumanAuthGetFaceVerifyAuthResult method
func (r *Mock) UnMockHumanAuthGetFaceVerifyAuthResult() {
	r.mockHumanAuthGetFaceVerifyAuthResult = nil
}

// GetFaceVerifyAuthResultReq ...
type GetFaceVerifyAuthResultReq struct {
	ReqOrderNo string  `query:"req_order_no" json:"-"` // 人脸识别单次唯一标识，由`tt.startFaceVerify`接口返回
	OpenID     *string `query:"open_id" json:"-"`      // 用户应用标识, 与employee_id二选其一
	EmployeeID *string `query:"employee_id" json:"-"`  // 用户租户标识, 与open_id二选其一
}

// getFaceVerifyAuthResultResp ...
type getFaceVerifyAuthResultResp struct {
	Code int64                        `json:"code,omitempty"` // 返回码，非0为失败
	Msg  string                       `json:"msg,omitempty"`  // 返回信息，返回码的描述
	Data *GetFaceVerifyAuthResultResp `json:"data,omitempty"` // 业务数据
}

// GetFaceVerifyAuthResultResp ...
type GetFaceVerifyAuthResultResp struct {
	AuthState     int64 `json:"auth_state,omitempty"`     // 认证结果, 0: 认证中, 1: 认证成功, -1: 认证失败
	AuthTimpstamp int64 `json:"auth_timpstamp,omitempty"` // 认证时间，unix 时间戳
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// UploadFaceVerifyImage
//
// 无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
// 无源人脸比对流程，开发者后台通过调用此接口将基准图片上传到飞书后台，做检测时的对比使用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/upload-facial-reference-image
func (r *HumanAuthService) UploadFaceVerifyImage(ctx context.Context, request *UploadFaceVerifyImageReq, options ...MethodOptionFunc) (*UploadFaceVerifyImageResp, *Response, error) {
	if r.cli.mock.mockHumanAuthUploadFaceVerifyImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#UploadFaceVerifyImage mock enable")
		return r.cli.mock.mockHumanAuthUploadFaceVerifyImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "UploadFaceVerifyImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/face_verify/v1/upload_face_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadFaceVerifyImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthUploadFaceVerifyImage mock HumanAuthUploadFaceVerifyImage method
func (r *Mock) MockHumanAuthUploadFaceVerifyImage(f func(ctx context.Context, request *UploadFaceVerifyImageReq, options ...MethodOptionFunc) (*UploadFaceVerifyImageResp, *Response, error)) {
	r.mockHumanAuthUploadFaceVerifyImage = f
}

// UnMockHumanAuthUploadFaceVerifyImage un-mock HumanAuthUploadFaceVerifyImage method
func (r *Mock) UnMockHumanAuthUploadFaceVerifyImage() {
	r.mockHumanAuthUploadFaceVerifyImage = nil
}

// UploadFaceVerifyImageReq ...
type UploadFaceVerifyImageReq struct {
	OpenID     *string   `query:"open_id" json:"-"`     // 用户应用标识, 与employee_id二选其一
	EmployeeID *string   `query:"employee_id" json:"-"` // 用户租户标识, 与open_id二选其一
	Image      io.Reader `json:"image,omitempty"`       // 带有头像的人脸照片
}

// uploadFaceVerifyImageResp ...
type uploadFaceVerifyImageResp struct {
	Code int64                      `json:"code,omitempty"` // 返回码，非0为失败
	Msg  string                     `json:"msg,omitempty"`  // 返回信息，返回码的描述
	Data *UploadFaceVerifyImageResp `json:"data,omitempty"` // 业务数据
}

// UploadFaceVerifyImageResp ...
type UploadFaceVerifyImageResp struct {
	FaceUid string `json:"face_uid,omitempty"` // 人脸图片用户Uid，需返回给应用小程序，作为小程序调起人脸识别接口的uid参数
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// CreateIdentity 该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。
//
// 实名认证接口会有计费管理，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/identity/create
func (r *HumanAuthService) CreateIdentity(ctx context.Context, request *CreateIdentityReq, options ...MethodOptionFunc) (*CreateIdentityResp, *Response, error) {
	if r.cli.mock.mockHumanAuthCreateIdentity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] HumanAuth#CreateIdentity mock enable")
		return r.cli.mock.mockHumanAuthCreateIdentity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "CreateIdentity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/human_authentication/v1/identities",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createIdentityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthCreateIdentity mock HumanAuthCreateIdentity method
func (r *Mock) MockHumanAuthCreateIdentity(f func(ctx context.Context, request *CreateIdentityReq, options ...MethodOptionFunc) (*CreateIdentityResp, *Response, error)) {
	r.mockHumanAuthCreateIdentity = f
}

// UnMockHumanAuthCreateIdentity un-mock HumanAuthCreateIdentity method
func (r *Mock) UnMockHumanAuthCreateIdentity() {
	r.mockHumanAuthCreateIdentity = nil
}

// CreateIdentityReq ...
type CreateIdentityReq struct {
	UserID       string  `query:"user_id" json:"-"`       // 用户的唯一标识（使用的ID类型见下一参数描述，不同ID类型的区别和获取，参考文档：[如何获得 User ID、Open ID 和 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get)）, 示例值: "ou_2eb5483cb377daa5054bc6f86e2089a5"
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户ID类型, 示例值: "open_id", 可选值有: `open_id`：用户的open id, `union_id`：用户的union id, `user_id`：用户的user id, 默认值: `open_id`, 当值为 `user_id`，字段权限要求: 获取用户 user ID
	IdentityName string  `json:"identity_name,omitempty"` // 姓名, 示例值: "张三"
	IdentityCode string  `json:"identity_code,omitempty"` // 身份证号, 示例值: "4xxxxxxxx"
	Mobile       *string `json:"mobile,omitempty"`        // 手机号, 示例值: "13xxxxxxx"
}

// createIdentityResp ...
type createIdentityResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *CreateIdentityResp `json:"data,omitempty"`
}

// CreateIdentityResp ...
type CreateIdentityResp struct {
	VerifyUid string `json:"verify_uid,omitempty"` // 用户绑定实名身份的uid
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// GetJssdkTicket
//
// 通过在你的网页中引入**飞书网页组件**，可在你的网页中直接使用飞书的功能。
// - 网页组件只适用于自建应用，暂不支持商店应用。
// - 网页组件适用于普通web网页，不建议在小程序webview中调用此组件
// ## 准备工作
// - 在开发者后台[创建一个企业自建应用](https://open.feishu.cn/document/home/introduction-to-custom-app-development/self-built-application-development-process)。
// - 引入组件库。在网页 html 中引入如下代码：
// ```html
// <script src="https://lf1-cdn-tos.bytescm.com/goofy/locl/lark/external_js_sdk/h5-js-sdk-1.0.11.js"></script>
// ```
// 若要使用成员卡片组件，SDK需要在`<body>`加载后引入。
// ## 鉴权流程
// ### 1. 获取 access_token
// - 不同的 token 代表了组件使用者的身份。user_access_token代表以用户身份鉴权，app_access_token代表以应用身份授权。
// - 成员名片组件仅支持以用户身份(user_access_token)鉴权
// - 云文档组件可以同时支持以用户身份(user_access_token)和应用身份(app_access_token)授权。但是以应用身份授权云文档时，访问量受 80 次/分钟限制，且组件不支持 “编辑”、“评论”、“点赞” 等功能
// - 开发者需要通过以下两种方式之一获取 token，再通过接口生成 ticket。
// -  方法一：获取用户身份。通过 [第三方网站免登](https://open.feishu.cn/document/ukTMukTMukTM/uETOwYjLxkDM24SM5AjN)获得 `user_access_token`
// - 方法二：获取应用身份。通过[服务端API](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token_internal)获得 `app_access_token`
// ### 2. 获取 jsapi_ticket
// 为了降低泄漏风险，这一步应当在你的服务端进行。
//
// doc: https://open.feishu.cn/document/uYjL24iN/uUDO3YjL1gzN24SN4cjN
func (r *JssdkService) GetJssdkTicket(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error) {
	if r.cli.mock.mockJssdkGetJssdkTicket != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Jssdk#GetJssdkTicket mock enable")
		return r.cli.mock.mockJssdkGetJssdkTicket(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Jssdk",
		API:                 "GetJssdkTicket",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/jssdk/ticket/get",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedAppAccessToken:  true,
		NeedUserAccessToken: true,
	}
	resp := new(getJssdkTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockJssdkGetJssdkTicket mock JssdkGetJssdkTicket method
func (r *Mock) MockJssdkGetJssdkTicket(f func(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error)) {
	r.mockJssdkGetJssdkTicket = f
}

// UnMockJssdkGetJssdkTicket un-mock JssdkGetJssdkTicket method
func (r *Mock) UnMockJssdkGetJssdkTicket() {
	r.mockJssdkGetJssdkTicket = nil
}

// GetJssdkTicketReq ...
type GetJssdkTicketReq struct{}

// getJssdkTicketResp ...
type getJssdkTicketResp struct {
	Code int64               `json:"code,omitempty"` // `返回码，非 0 表示失败`
	Msg  string              `json:"msg,omitempty"`  // `对返回码的文本描述`
	Data *GetJssdkTicketResp `json:"data,omitempty"` // `返回内容`
}

// GetJssdkTicketResp ...
type GetJssdkTicketResp struct {
	ExpireIn int64  `json:"expire_in,omitempty"` // `jsapi_ticket 的有效时间`
	Ticket   string `json:"ticket,omitempty"`    // `jsapi_ticket`
}

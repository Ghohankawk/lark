// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"github.com/chyroc/lark/internal"
)

// OpenCalender 跳转并打开日历，打开界面为上一次离开日历时的视图。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-calender/open-a-calender
func (r *AppLinkService) OpenCalender(req *OpenCalenderReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/calendar/open", req)
}

type OpenCalenderReq struct{}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenCalenderAccount 打开第三方日历账户管理页，方便用户添加或导入谷歌/Exchange 日历。移动端打开页面，PC端打开弹层。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-calender/open
func (r *AppLinkService) OpenCalenderAccount(req *OpenCalenderAccountReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/calendar/account", req)
}

type OpenCalenderAccountReq struct{}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenCalenderEventCreate 跳转日历 tab 并进入日程创建页面，用户可新建日程。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-calender/open-the-schedule-creation-page
func (r *AppLinkService) OpenCalenderEventCreate(req *OpenCalenderEventCreateReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/calendar/event/create", req)
}

type OpenCalenderEventCreateReq struct {
	StartTime *string `json:"startTime,omitempty"` // 开始日期，{unixTime}格式
	EndTime   *string `json:"endTime,omitempty"`   // 结束日期，{unixTime}格式
	Summary   *string `json:"summary,omitempty"`   // 日程主题，中文可使用encodeURIComponent方法生成
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenCalenderView 打开日历tab，并支持定义跳转到具体视图和具体日期。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-calender/open-a-calendar-and-support-to-define-view-and-date
func (r *AppLinkService) OpenCalenderView(req *OpenCalenderViewReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/calendar/view", req)
}

type OpenCalenderViewReq struct {
	Type *string `json:"type,omitempty"` // 视图类型，枚举值包括：<br> `day`：日视图 <br>`three_day`：三日视图，仅移动端支持<br> `week`：周视图，仅PC端支持 <br>  `month`：月视图<br> `meeting`：会议室视图，仅PC端支持 <br> `list`：列表视图，仅移动端支持
	Date *string `json:"date,omitempty"` // 日期，{unixTime}格式
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenChat 打开一个聊天页面，单聊会话或群聊会话（仅能打开用户已加入的单聊或群聊会话，不会自动进入未加入的群组）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-a-chat-page
func (r *AppLinkService) OpenChat(req *OpenChatReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/chat/open", req)
}

type OpenChatReq struct {
	OpenID     string `json:"openId,omitempty"`     // 用户 openId
	OpenChatID string `json:"openChatId,omitempty"` // 会话ID，包括单聊会话和群聊会话。示例：oc_41e7bdf4877cfc316136f4ccf6c32613
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenDocs 打开云文档（docs）。使用外部浏览器打开文档时，提供入口从飞书中打开。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-docs
func (r *AppLinkService) OpenDocs(req *OpenDocsReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/docs/open", req)
}

type OpenDocsReq struct {
	URL string `json:"URL,omitempty"` // 要打开的云文档URL
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenLark 唤起飞书客户端
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-lark
func (r *AppLinkService) OpenLark(req *OpenLarkReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/op/open", req)
}

type OpenLarkReq struct{}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenMiniProgram 打开一个小程序或者小程序中的一个页面
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-a-gadget
func (r *AppLinkService) OpenMiniProgram(req *OpenMiniProgramReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/mini_program/open", req)
}

type OpenMiniProgramReq struct {
	AppID       string  `json:"appId,omitempty"`        // 小程序 appId(可从「开发者后台-凭证与基础信息」获取)
	Mode        *string `json:"mode,omitempty"`         // PC小程序启动模式，枚举值包括：<br>`sidebar-semi`：聊天的侧边栏打开<br>`appCenter`：工作台中打开<br>`window`：独立大窗口打开<br>`window-semi`：独立小窗口打开，3.33版本开始支持此模式
	Path        *string `json:"path,omitempty"`         // 需要跳转的页面路径，路径后可以带参数。也可以使用 path_android、path_ios、path_pc 参数对不同的客户端指定不同的path
	PathAndroid *string `json:"path_android,omitempty"` // 同 path 参数，Android 端会优先使用该参数，如果该参数不存在，则会使用 path 参数
	PathIos     *string `json:"path_ios,omitempty"`     // 同 path 参数，iOS 端会优先使用该参数，如果该参数不存在，则会使用 path 参数
	PathPc      *string `json:"path_pc,omitempty"`      // 同 path 参数，PC 端会优先使用该参数，如果该参数不存在，则会使用 path 参数
	MinLkVer    *string `json:"min_lk_ver,omitempty"`   // 指定 AppLink 协议能够兼容的最小飞书版本，使用三位版本号 x.y.z。如果当前飞书版本号小于min_lk_ver，打开该 AppLink 会显示为兼容页面
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenSSOLogin 在飞书客户端中打开租户在admin后台配置的SSO登录页
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-the-sso-login-page
func (r *AppLinkService) OpenSSOLogin(req *OpenSSOLoginReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/passport/sso_login", req)
}

type OpenSSOLoginReq struct {
	SSODomain  string `json:"sso_domain,omitempty"`  // 租户的域名，填写的是租户在admin后台配置的租户域名信息。当在admin后台改动租户的域名时，需要同步修改applink该参数值
	TenantName string `json:"tenant_name,omitempty"` // 租户名，用于在切换租户时，客户端展示即将登录到的租户名称，一般填写公司名即可
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenWebApp 打开一个已安装的H5应用
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-an-h5-app
func (r *AppLinkService) OpenWebApp(req *OpenWebAppReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/web_app/open", req)
}

type OpenWebAppReq struct {
	AppID       string  `json:"appId,omitempty"`        // H5应用的 appId(可从「开发者后台-凭证与基础信息」获取)
	Mode        *string `json:"mode,omitempty"`         // 打开H5应用的容器模式，枚举值包括<br> `appCenter`：在工作台打开，3.20版本开始支持（缺省值） <br> `window`：在独立窗口打开，3.20版本开始支持 <br> `sidebar`：在侧边栏打开，3.40版本开始支持
	Path        *string `json:"path,omitempty"`         // 访问H5应用的具体某个页面，path参数将替换H5应用URL的path部分（注意：path中不应该出现#和?字符，否则会导致最终的H5页面URL结构异常） <br>也可以使用 path_android、path_ios、path_pc 参数对不同的客户端指定不同的path <br>3.20版本开始支持
	PathAndroid *string `json:"path_android,omitempty"` // 同 path 参数，Android 端会优先使用该参数，如果该参数不存在，则会使用 path 参数。 <br>3.20版本开始支持
	PathIos     *string `json:"path_ios,omitempty"`     // 同 path 参数，iOS 端会优先使用该参数，如果该参数不存在，则会使用 path 参数 <br>3.20版本开始支持
	PathPc      *string `json:"path_pc,omitempty"`      // 同 path 参数，PC 端会优先使用该参数，如果该参数不存在，则会使用 path 参数 <br>3.20版本开始支持
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// OpenWebURL 用户在PC端点击这类applink，不必跳转外部浏览器，可以直接在聊天的侧边栏、或飞书的独立窗口中打开指定的网页。,可以配置这类applink在消息卡片的“查看详情”跳转上，使用户连贯地消费消息中的详情内容。,,![image.png](//sf3-cn.feishucdn.com/obj/open-platform-opendoc/814bd45433ea80bba3faf7f421050ea4_dMDgWVZDOl.png)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-the-web-view-in-feishu-to-access-the-specified-url
func (r *AppLinkService) OpenWebURL(req *OpenWebURLReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/web_url/open", req)
}

type OpenWebURLReq struct {
	URL    string  `json:"url,omitempty"`    // 指定需要在客户端内打开的具体链接，需要执行encodeURIComponent，4.2+版本支持lark协议
	Mode   string  `json:"mode,omitempty"`   // 打开的容器模式，枚举值包括：<br>`sidebar-semi` 在侧边栏打开；<br>`window` 在独立窗口打开
	Height *string `json:"height,omitempty"` // 自定义独立窗口高度
	Width  *string `json:"width,omitempty"`  // 自定义独立窗口宽度
}
package types

// -----------------------------------

// 成员管理

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加

type CreateMemberRequest struct {
	Nickname string   `json:"nickname" binding:"required,min=4,max=20"`          // required，不小于 4 位 不超过 20 位
	Username string   `json:"username" binding:"required,alpha,min=8,max=20"`    // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   `json:"password" binding:"required,alphanum,min=8,max=20"` // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType `json:"userType" binding:"required,min=1,max=3"`           // required, 枚举值
}

type CreateMemberResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		UserID string `json:"userID"` // int64 范围
	} `json:"data"`
}

// 获取成员信息

type GetMemberRequest struct {
	UserID string `json:"userID" binding:"required"`
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code ErrNo   `json:"code"`
	Data TMember `json:"data"`
}

// 批量获取成员信息

type GetMemberListRequest struct {
	Offset int `json:"offset" binding:"required"`
	Limit  int `json:"limit" binding:"required"`
}

type GetMemberListResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		MemberList []TMember `json:"memberList"`
	} `json:"data"`
}

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string `json:"userID" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type UpdateMemberResponse struct {
	Code ErrNo `json:"code"`
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string `json:"userID" binding:"required"`
}

type DeleteMemberResponse struct {
	Code ErrNo `json:"code"`
}

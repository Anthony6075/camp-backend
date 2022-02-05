package types

// -----------------------------------

// 成员管理

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加

type CreateMemberRequest struct {
	Nickname string   // required，不小于 4 位 不超过 20 位
	Username string   // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType // required, 枚举值
}

type CreateMemberResponse struct {
	Code ErrNo
	Data struct {
		UserID string // int64 范围
	}
}

// 获取成员信息

type GetMemberRequest struct {
	UserID string
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code ErrNo
	Data TMember
}

// 批量获取成员信息

type GetMemberListRequest struct {
	Offset int
	Limit  int
}

type GetMemberListResponse struct {
	Code ErrNo
	Data struct {
		MemberList []TMember
	}
}

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string
	Nickname string
}

type UpdateMemberResponse struct {
	Code ErrNo
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string
}

type DeleteMemberResponse struct {
	Code ErrNo
}

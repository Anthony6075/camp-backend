package types

// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

type TMember struct {
	UserID    string   `json:"userID" gorm:"primaryKey"`
	Nickname  string   `json:"nickname"`
	Username  string   `json:"username" gorm:"unique;not null"`
	Password  string   `json:"-" gorm:"not null"`
	UserType  UserType `json:"userType" gorm:"not null"`
	IsDeleted bool     `json:"-" gorm:"not null;default:false"`
}

type TCourse struct {
	CourseID  string
	Name      string
	TeacherID string
}

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

func (u UserType) String() string {
	switch u {
	case Admin:
		return "Admin"
	case Student:
		return "Student"
	case Teacher:
		return "Teacher"
	default:
		return "Unknown"
	}
}

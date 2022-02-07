package types

// -------------------------------------
// 排课

// 创建课程
// Method: Post
type CreateCourseRequest struct {
	Name string `json:"name" binding:"required"`
	Cap  int    `json:"cap" binding:"required"`
}

type CreateCourseResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		CourseID string `json:"courseID"`
	} `json:"data"`
}

// 获取课程
// Method: Get
type GetCourseRequest struct {
	CourseID string `json:"courseID" binding:"required"`
}

type GetCourseResponse struct {
	Code ErrNo   `json:"code"`
	Data TCourse `json:"data"`
}

// 老师绑定课程
// Method： Post
// 注：这里的 teacherID 不需要做已落库校验
// 一个老师可以绑定多个课程 , 不过，一个课程只能绑定在一个老师下面
type BindCourseRequest struct {
	CourseID  string `json:"courseID" binding:"required"`
	TeacherID string `json:"teacherID" binding:"required"`
}

type BindCourseResponse struct {
	Code ErrNo `json:"code"`
}

// 老师解绑课程
// Method： Post
type UnbindCourseRequest struct {
	CourseID  string `json:"courseID" binding:"required"`
	TeacherID string `json:"teacherID" binding:"required"`
}

type UnbindCourseResponse struct {
	Code ErrNo `json:"code"`
}

// 获取老师下所有课程
// Method：Get
type GetTeacherCourseRequest struct {
	TeacherID string `json:"teacherID" binding:"required"`
}

type GetTeacherCourseResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		CourseList []TCourse `json:"courseList"`
	} `json:"data"`
}

// 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程
// Method： Post
type ScheduleCourseRequest struct {
	TeacherCourseRelationShip map[string][]string // key 为 teacherID , val 为老师期望绑定的课程 courseID 数组
}

type ScheduleCourseResponse struct {
	Code ErrNo
	Data map[string]string // key 为 teacherID , val 为老师最终绑定的课程 courseID
}

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// 课程已满返回 CourseNotAvailable

type BookCourseResponse struct {
	Code ErrNo
}

type GetStudentCourseRequest struct {
	StudentID string
}

type GetStudentCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseList []TCourse
	}
}

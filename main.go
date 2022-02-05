package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	setupDatasource()

	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", login)
		auth.POST("/logout", logout)
		auth.GET("/whoami", whoami)
	}

	member := r.Group("/member")
	{
		member.POST("/create", createMember)
		member.GET("", getMember)
		member.GET("/list", getMemberList)
		member.POST("/update", updateMember)
		member.POST("/delete", deleteMember)
	}

	course := r.Group("/course")
	{
		course.POST("/create", createCourse)
		course.GET("/get", getCourse)
		course.POST("/schedule", scheduleCourse)
	}

	courseTeacher := r.Group("/teacher")
	{
		courseTeacher.POST("/bind_course", bindCourse)
		courseTeacher.POST("/unbind_course", unbindCourse)
		courseTeacher.GET("/get_course", getTeacherCourse)
	}

	student := r.Group("/student")
	{
		student.POST("/book_course", bookCourse)
		student.GET("/course", getStudentCourse)
	}

	r.GET("hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	return r
}

package main

import (
	"camp-backend/handler"
	"camp-backend/initial"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	initial.SetupDatasource()
	initial.SetupRedis()

	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	d := gin.Default()

	d.GET("/test1", func(c *gin.Context) {
		p := c.Query("p")
		if p == "1" {
			time.Sleep(10 * time.Second)
			c.JSON(http.StatusOK, gin.H{"111": "OK"})
		} else {
			time.Sleep(2 * time.Second)
			c.JSON(http.StatusOK, gin.H{"other": "OK"})
		}
	})

	r := d.Group("/api/v1")

	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.GET("/whoami", handler.Whoami)
	}

	member := r.Group("/member")
	{
		member.POST("/create", handler.CreateMember)
		member.GET("", handler.GetMember)
		member.GET("/list", handler.GetMemberList)
		member.POST("/update", handler.UpdateMember)
		member.POST("/delete", handler.DeleteMember)
	}

	course := r.Group("/course")
	{
		course.POST("/create", handler.CreateCourse)
		course.GET("/get", handler.GetCourse)
		course.POST("/schedule", handler.ScheduleCourse)
	}

	courseTeacher := r.Group("/teacher")
	{
		courseTeacher.POST("/bind_course", handler.BindCourse)
		courseTeacher.POST("/unbind_course", handler.UnbindCourse)
		courseTeacher.GET("/get_course", handler.GetTeacherCourse)
	}

	student := r.Group("/student")
	{
		student.POST("/book_course", handler.BookCourse)
		student.GET("/course", handler.GetStudentCourse)
	}

	r.GET("hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	return d
}

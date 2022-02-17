package handler

import (
	"camp-backend/initial"
	"camp-backend/types"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func BookCourse(c *gin.Context) {

}

func GetStudentCourse(c *gin.Context) {
	request := new(types.GetStudentCourseRequest)
	response := new(types.GetStudentCourseResponse)
	fmt.Printf("%+v\n", request)

	request.StudentID = c.Query("StudentID")
	if request.StudentID == "" {
		response.Code = types.ParamInvalid
		c.JSON(http.StatusOK, response)
		fmt.Printf("%+v \n", response.Code)

		return
	}

	theUser := new(types.TMember)
	err := initial.Db.First(theUser, request.StudentID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || theUser.UserType.String() != "Student" {
		response.Code = types.StudentNotExisted
		c.JSON(http.StatusOK, response)
		fmt.Printf("%+v \n", response.Code)

		return
	}

	courses := make([]types.TCourse, 0)
	if err := initial.Db.Model(theUser).Association("LearnCourses").Find(&courses); err != nil {
		response.Code = types.UnknownError
		c.JSON(http.StatusOK, response)
		fmt.Printf("%+v \n", response.Code)

		return
	}

	if len(courses) == 0 {
		response.Code = types.StudentHasNoCourse
	} else {
		response.Code = types.OK
	}
	response.Data.CourseList = courses
	c.JSON(http.StatusOK, response)
	fmt.Printf("%+v \n", response.Code)

}

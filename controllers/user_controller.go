package controllers

import (
	"fmt"
	"github.com/AhmedSafwat1/micro_service_go-api/domain/users"
	"github.com/AhmedSafwat1/micro_service_go-api/services"
	"github.com/AhmedSafwat1/micro_service_go-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId , err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	fmt.Println(userId)
	if err != nil {
		err := errors.NewBadRequest("invald user id")
		c.JSON(err.Code, err)
		return
	}
	result ,myerr := services.GetUser(userId)
	//fmt.Println(result, myerr)
	if myerr != nil{
		c.JSON(myerr.Code, myerr)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context)  {
	var user users.User




	if err:= c.ShouldBindJSON(&user); err != nil{
		//handl error
		rest := errors.NewBadRequest("invalid json body")
		fmt.Println("error : ", err)
		c.JSON(rest.Code, rest)
		return
	}

	result ,err := services.CreateUser(user);
	if err != nil{
		c.JSON(err.Code, err)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

func Search(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "not implement")
}



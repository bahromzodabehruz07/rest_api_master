package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rest_api_master/domains/error"
	"github.com/rest_api_master/domains/response"
	"github.com/rest_api_master/domains/user"
	"github.com/rest_api_master/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(con *gin.Context) {

	//getting request body
	body, err := ioutil.ReadAll(con.Request.Body)

	//empty models for sending customresponse
	user := user.User{}
	response := response.Response{}

	//checking for reading body
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Bad date"
		response.Payload = nil
		con.JSON(http.StatusBadRequest, response)
		return
	}
	//parsing body to user struct
	parserErr := json.Unmarshal(body, &user)

	//checking for parsing result
	if parserErr != nil {
		//handle bad request error
		response.Status = http.StatusBadRequest
		response.Message = "Bad date"
		response.Payload = nil
		con.JSON(http.StatusBadRequest, response)
		return
	}

	//var validEmail = utils.IsEmailValid(user.Email)
	////checking valid email
	//
	//if !validEmail {
	//	response.Status = http.StatusBadRequest
	//	response.Message = "Please send valid email"
	//	response.Payload = nil
	//	con.JSON(http.StatusBadRequest, response)
	//	return
	//}

	//if parsing is ok
	insertRes := user.SaveUser()
	//checking insert result
	if insertRes != nil {
		//if insert query return error we will send response error to front
		con.JSON(http.StatusNotImplemented, insertRes)
		return
	}

	con.JSON(http.StatusOK, user)
}

func GetAllUsers(con *gin.Context) {
	result, err := service.GetAllUsers()
	if err != nil {
		con.JSON(http.StatusNotFound, err)

	}
	con.JSON(http.StatusOK, result)

}

func UpdateUser(con *gin.Context) {
	userId, parseErr := strconv.ParseInt(con.Param("userId"), 10, 64)
	if parseErr != nil {
		err_response := error.Error{
			Message: "Bad data",
			Status:  http.StatusBadRequest,
			Error:   parseErr.Error(),
		}
		con.JSON(http.StatusBadRequest, err_response)
	}
	fmt.Println(userId)
	var user user.User
	if err := con.BindJSON(&user); err != nil {
		err_response := error.Error{
			Message: "Bad data",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		con.JSON(http.StatusBadRequest, err_response)
	}
	user.Id = userId
	fmt.Println(user)
	res := user.UpdateUser()
	if res != nil {
		panic(res)
	}

	con.JSON(http.StatusOK, "Called UpdateUser Method")
}

func DeleteUser(con *gin.Context) {
	userId, parseErr := strconv.ParseInt(con.Param("userId"), 10, 64)
	if parseErr != nil {
		err_response := error.Error{
			Message: "Bad data",
			Status:  http.StatusBadRequest,
			Error:   "Please input valid id",
		}
		con.JSON(http.StatusBadRequest, err_response)
		return
	}
	user := user.User{Id: userId}
	err := user.Delete()
	if err != nil {
		con.JSON(http.StatusNotFound, err)
		return
	}

	con.JSON(http.StatusOK, "Called DeleteUser Method")
}

func GetUserById(con *gin.Context) {
	userId, parseErr := strconv.ParseInt(con.Param("userId"), 10, 64)
	if parseErr != nil {
		err_response := error.Error{
			Message: "Bad data",
			Status:  http.StatusBadRequest,
			Error:   "Please input valid id",
		}
		con.JSON(http.StatusBadRequest, err_response)
		return
	}
	user := user.User{Id: userId}
	err := user.GetUser()
	if err != nil {
		con.JSON(http.StatusNotFound, err)
		return
	}
	con.JSON(http.StatusOK, user)

}

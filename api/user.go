package api

import (
	"resource-plan-improvement/config"
	"resource-plan-improvement/entity"
	"resource-plan-improvement/service"
	"resource-plan-improvement/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	var err error
	var signin entity.Signin
	if err = c.BindJSON(&signin); err != nil {
		FailWithMsg(c, "Error parameters")
		return
	}
	if err = util.Authenticate(config.Conf.Ldap.URL, signin.Username, signin.Password); err != nil {
		FailWithMsg(c, "wrong username or password")
		return
	}

	var user entity.User
	if user, err = entity.FindUserBySAMAccount(signin.Username); err != nil {
		FailWithMsg(c, err.Error())
		return
	}

	var token string
	if token, err = service.GenerateToken(user.Id); err != nil {
		AbortInternalServerError(c, err)
		return
	}
	user.Token = token

	OkWithData(c, user)
}

func AutoSignin(c *gin.Context) {
	var err error
	var autoSignin entity.AutoSignin
	if err = c.BindHeader(&autoSignin); err != nil {
		FailWithMsg(c, "Error parameters")
		return
	}

	var userId uint
	if userId, err = service.GetUserIdFromToken(autoSignin.Token); err != nil {
		FailWithMsg(c, err.Error())
		return
	}

	var user entity.User
	if user, err = entity.FindUserById(userId); err != nil {
		FailWithMsg(c, err.Error())
		return
	}

	user.Token = autoSignin.Token

	OkWithData(c, user)
}

func Signup(c *gin.Context) {
	AbortNotImplementedMethod(c)
}

func Signout(c *gin.Context) {
	AbortNotImplementedMethod(c)
}

func GetAllUsers(c *gin.Context) {
	if users, err := entity.FindAllUsers(); err != nil {
		FailWithErr(c, err)
	} else {
		OkWithData(c, users)
	}
}

func GetUserById(c *gin.Context) {
	var userId int
	var err error
	if userId, err = strconv.Atoi(c.Param("id")); err != nil {
		AbortBadRequest(c, err)
		return
	}
	var user entity.User
	if user, err = entity.FindUserById(uint(userId)); err != nil {
		AbortFindFailed(c, err)
		return
	}
	OkWithData(c, user)
}

func DeleteUserById(c *gin.Context) {
	var userId int
	var err error
	if userId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = service.DeleteUser(uint(userId)); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}

func SaveUser(c *gin.Context) {
	var err error
	var user entity.User
	if err = c.BindJSON(&user); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	if err = user.Save(); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	OkWithData(c, user)
}

func GetUsersUnderLeader(c *gin.Context) {
	userId := c.GetUint(config.CTX_KEY_USER_ID)
	if users, err := service.GetUsersUnderLeader(userId); err != nil {
		FailWithMsg(c, err.Error())
	} else {
		OkWithData(c, users)
	}
}

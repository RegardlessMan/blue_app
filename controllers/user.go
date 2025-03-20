package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

func SignUpHandler(c *gin.Context) {
	// 1、获取参数和校验
	p := &models.ParamSignUp{}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	//2、业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "sign up failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "sign up success",
	})
}

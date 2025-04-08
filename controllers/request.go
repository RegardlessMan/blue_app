package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

const CtxUserIDKey = "userID"

// GetCurrentUserID 获取当前登录用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		ResponseError(c, CodeNeedLogin)
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		zap.L().Error("GetCurrentUserID type assertion failed", zap.Any("c.Get(CtxUserIDKey)", uid))
		ResponseError(c, CodeServerBusy)
		return
	}
	return userID, nil
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}

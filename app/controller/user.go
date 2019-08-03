package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/library"
)

func Login(c *gin.Context) {

	//获取参数
	var param struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	c.Bind(&param)
	uid := library.GetMd5(param.Username)
	//生成token
	token, err := library.GenerateToken(uid, param.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    201,
			"message": err,
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"token":   token,
			"uid":     uid,
		})
	}

}

package login

import (
	"eassy/core/jwt"
	"eassy/model"
	"eassy/util"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Pwd   string `form:"password" json:"password" xml:"password" binding:"required"`
}

type RegForm struct {
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Pwd   string `form:"password" json:"password" xml:"password" binding:"required"`
}

type VerifyForm struct {
	Token string `form:"token" json:"token" xml:"token" binding:"required"`
	//AccountId   string `form:"accountId" json:"accountId" xml:"accountId" binding:"required"`
}

func loginHandler(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		util.Failed(ctx, err.Error(), nil)
		return
	}
	acc := model.GetAccountByPhone(form.Phone)
	if acc.ID == 0 || !util.BCryptVerify([]byte(form.Pwd), []byte(acc.Pwd)) {
		util.Failed(ctx, "登录失败!", nil)
		return
	}
	token, _ := jwt.ReleaseToken(acc.ID)
	util.Success(ctx, "登陆成功！", gin.H{
		"token":     token,
		"accountId": acc.ID,
	})
}

func registerHandler(ctx *gin.Context) {
	var form RegForm
	if err := ctx.ShouldBind(&form); err != nil {
		util.Failed(ctx, err.Error(), nil)
		return
	}
	pwd, _ := util.BCrypt(form.Pwd)
	err := model.RegAccount(form.Phone, pwd)
	if err != nil {
		util.Failed(ctx, "注册失败,err: "+err.Error(), nil)
		return
	}
	util.Success(ctx, "注册成功!", nil)
}

func verifyHandler(ctx *gin.Context) {
	var form VerifyForm
	if err := ctx.ShouldBind(&form); err != nil {
		util.Failed(ctx, err.Error(), nil)
		return
	}
	token, claims, err := jwt.ParseToken(form.Token)
	if err != nil || !token.Valid {
		util.Failed(ctx, "token验证失败！", nil)
		ctx.Abort()
		return
	}
	accId := claims.UserId
	//验证accId
	if !model.IsExist(accId) {
		util.Failed(ctx, "无效的用户！", nil)
		ctx.Abort()
		return
	}
	util.Success(ctx, "验证成功", gin.H{
		"accId": accId,
	})
}

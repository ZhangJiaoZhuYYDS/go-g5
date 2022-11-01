// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package middleware

import (
	"b5gocmf/admin/controller/common"
	"b5gocmf/admin/lib"
	"b5gocmf/admin/services"
	"b5gocmf/utils/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LoginAdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		publicC := common.NewPublicController()
		publicUrl := publicC.ParseUrl("", true)
		path := ctx.Request.URL.Path
		if strings.Index(path, publicUrl) == 0 {
			return
		}
		data, err := services.CheckLoginCookie(ctx)
		if err != nil || data.Id == "" {
			loginUrl := publicC.ParseUrl("login", true)
			if tool.IsRender(ctx) {
				ctx.Redirect(http.StatusFound, loginUrl)
			} else {
				publicC.Error(ctx, "请先登录", nil, lib.RNoLogin, loginUrl)
			}
			ctx.Abort()
			return
		}else{
			//往中间件中设置登录信息
			ctx.Set("_login_",data)
		}
	}
}

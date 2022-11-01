// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package common

import (
	"b5gocmf/admin/lib"
	"b5gocmf/admin/services"
	"b5gocmf/utils/types"
	"github.com/gin-gonic/gin"
	"html/template"
)

func (c *IndexController) Route(engine *gin.Engine,group *gin.RouterGroup) {
	group.GET(c.Dispatch("index",false, c.Index))
	group.GET(c.Dispatch("home",false, c.Home))
}

type IndexController struct {
	lib.Controller
}

// NewIndexController 创建控制并初始化参数
func NewIndexController() *IndexController {
	c := &IndexController{}
	c.Group = ""
	c.Id = "index"
	return c
}

func (c *IndexController) Index(ctx *gin.Context) {
	userInfo := make(map[string]string)
	userInfo["nick_name"] = "小李"
	userInfo["struct"] = "冰舞网络"
	list := services.AdminMenuShowList(ctx)
	html := ""
	if list != nil && len(list)>0 {
		html = services.AdminMenuToHtml(list,0)
	}
	c.Render(ctx, "index", gin.H{"user": userInfo,"menuHtml":&types.HtmlShow{Html: template.HTML(html)}})
}

func (c *IndexController) Home(ctx *gin.Context) {
	c.Render(ctx, "home", gin.H{})
}

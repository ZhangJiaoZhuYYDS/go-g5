// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package demo

import (
	"b5gocmf/admin/lib"
	"b5gocmf/admin/services"
	. "b5gocmf/common/models/demo"
	"github.com/gin-gonic/gin"
)

// Route 定义该控制器的路由
func (c *TestInfoController) Route(engine *gin.Engine,group *gin.RouterGroup) {
    group.GET(c.Dispatch("index", false, c.Index))
    group.POST(c.Dispatch("index", false, c.FindList))
    group.GET(c.Dispatch("add", false, c.Add))
    group.POST(c.Dispatch("add", false, c.AddSave))
    group.GET(c.Dispatch("edit", false, c.Edit))
    group.POST(c.Dispatch("edit", false, c.EditSave))
    group.POST(c.Dispatch("drop", false, c.Drop))
    group.POST(c.Dispatch("drop_all", false, c.DropAll))
}

type TestInfoController struct {
	lib.Controller
}

// NewTestInfoController 创建控制并初始化参数
func NewTestInfoController() *TestInfoController {
	c := &TestInfoController{}
	c.Id = "test_info"
	c.Group = "demo"
	c.IModel = NewTestInfoModel()
	return c
}

// FindList 获取列表json
func (c *TestInfoController) FindList(ctx *gin.Context) {
	//获取数据权限的where语句和占位替换args
	params, args := services.NewDataScopeFilterByCtx(ctx).GetQueryParams("struct_id", "user_id")
	c.GetIndex(ctx, NewTestInfoModel().NewSlice(), params, args, nil)
}

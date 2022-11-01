// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package core

// ** 全局参数 ** //

import (
	"b5gocmf/config"
)

var (
	G_CONFIG    *config.Config //配置文件
	G_DB        *B5Db
	G_Redis		*B5Redis
	G_TIME      string = "2006-01-02 15:04:05"
	G_DATE      string = "2006-01-02"
	G_GENID     *Node
	G_Validate  *ValidateCtx
	G_STRUCT_ID string         = "100"   //根组织ID
	G_ROLE_ID   string         = "1"     //超管角色ID
	G_ADMIN_ID  string         = "10000" //超管ID
	G_DATA      = make(map[string]any,0)           //全局缓存数据，在启动时存放一些数据
)

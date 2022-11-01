// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package services

import (
	"b5gocmf/common/services/system"
	"b5gocmf/utils/tool"
	"b5gocmf/utils/types"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"sync"
)

// DataScopeTypeList 数据权限类型
func DataScopeTypeList() []types.KeyVal {
	return []types.KeyVal{{Key: "1", Value: "全部数据权限"}, {Key: "2", Value: "本部门及以下数据权限"}, {Key: "4", Value: "本部门数据权限"}, {Key: "8", Value: "自定数据权限"}, {Key: "16", Value: "仅本人数据权限"}}
}

type DataScopeFilter struct {
	LoginData *LoginData
	powerData *dataScopeUserPower
	sync.Mutex
}

func NewDataScopeFilterByCtx(ctx *gin.Context) *DataScopeFilter {
	return &DataScopeFilter{powerData:nil,LoginData:GetLoginByCtx(ctx)}
}

type dataScopeUserPower struct {
	NoPower   bool     // true 无任何权限
	AllPower  bool     // true 所有权限
	StructIds []string //
	IsUser    bool     //true代表自己的数据
}

// GetQueryParams 生成查询where条件 和 占位替代数组
func (ds *DataScopeFilter) GetQueryParams(structField string, userField string) (whereStr string, args []any) {
	whereStr = " 1 = 0 "
	args = make([]any, 0)

	if structField == "" && userField == "" { //字段都存在则默认无权限
		return
	}
	powerData := ds.getPowerData()
	if powerData.NoPower { //无权限
		return
	}
	if powerData.AllPower { //所有权限
		whereStr = ""
		return
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser { //未查询到组织及非个人 无权限
		return
	}

	if structField != "" && userField != "" { //判断组织和个人
		if len(powerData.StructIds) > 0 && powerData.IsUser {
			whereStr = " ( `" + userField + "` = ? OR `" + structField + "` in (?) ) "
			args = append(args, ds.LoginData.Id)
			args = append(args, strings.Join(powerData.StructIds, ","))
		} else if len(powerData.StructIds) > 0 {
			whereStr = " `" + structField + "` in (?) "
			args = append(args, strings.Join(powerData.StructIds, ","))
		} else if powerData.IsUser {
			whereStr = " `" + userField + "` = ? "
			args = append(args, ds.LoginData.Id)
		}
	} else if structField != "" {
		if len(powerData.StructIds) > 0 {
			whereStr = " `" + structField + "` in (?) "
			args = append(args, strings.Join(powerData.StructIds, ","))
		}
	} else if userField != "" {
		if powerData.IsUser {
			whereStr = " `" + userField + "` = ? "
			args = append(args, ds.LoginData.Id)
		}
	}
	return
}

//CheckByFiled 根据传入的组织ID和用户ID 判断是否有权限
func (ds *DataScopeFilter) CheckByFiled(structId string,userId string) bool {
	if structId == "" && userId == "" {
		return false
	}
	powerData := ds.getPowerData()
	if powerData.NoPower { //无权限
		return false
	}
	if powerData.AllPower { //所有权限
		return true
	}
	if len(powerData.StructIds) == 0 && !powerData.IsUser { //未查询到组织及非个人 无权限
		return false
	}
	if structId != "" && userId != "" { //判断组织和个人
		if len(powerData.StructIds) > 0 && powerData.IsUser {
			if userId == ds.LoginData.Id {
				return true
			}
			if tool.InArray(structId,powerData.StructIds) {
				return true
			}
		} else if len(powerData.StructIds) > 0 {
			if tool.InArray(structId,powerData.StructIds) {
				return true
			}
		} else if powerData.IsUser {
			if userId == ds.LoginData.Id {
				return true
			}
		}
	} else if structId != "" {
		if len(powerData.StructIds) > 0 {
			if tool.InArray(structId,powerData.StructIds) {
				return true
			}
		}
	} else if userId != "" {
		if powerData.IsUser {
			if userId == ds.LoginData.Id {
				return true
			}
		}
	}
	return false
}

//getPowerData 获取根据parseDataScopePower解析的权限
func (ds *DataScopeFilter) getPowerData() *dataScopeUserPower {
	ds.Lock()
	if ds.powerData == nil{
		ds.powerData = ds.parseDataScopePower()
	}
	ds.Unlock()
	return ds.powerData
}
// parseDataScopePower 根据登录信息解析权限 生成dataScopeUserParse
func (ds *DataScopeFilter) parseDataScopePower() *dataScopeUserPower {
	res := &dataScopeUserPower{}

	if ds.LoginData == nil {
		res.NoPower = true //无权限
		return res
	}
	userId, _ := strconv.Atoi(ds.LoginData.Id)
	if userId < 1 {
		res.NoPower = true //无权限
		return res
	}
	if ds.LoginData.IsAdmin == "1" {
		res.AllPower = true //超管返回全部权限
		return res
	}

	dataScope, _ := strconv.Atoi(ds.LoginData.DataScope)
	if dataScope < 1 {
		res.NoPower = true //无权限
		return res
	}

	if (31 & dataScope) == 0 {
		res.NoPower = true //无效的键值
		return res
	}

	structIdInt, _ := strconv.Atoi(ds.LoginData.StructId)
	if structIdInt < 1 {
		res.NoPower = true //无组织 返回无权限
		return res
	}

	if len(ds.LoginData.RoleList) < 1 {
		res.NoPower = true //无角色返回无权限
		return res
	}

	//权限范围的组织ID
	structIds := make([]string, 0)

	if 1&dataScope == 1 { //全部数据权限
		res.AllPower = true
		return res
	}

	if 2&dataScope == 2 { //本部门及以下数据权限
		structIds = append(structIds, ds.LoginData.StructId)
		childList := system.NewStructService().GetChildAllIdList(ds.LoginData.StructId)
		if len(childList) > 0 {
			structIds = append(structIds, childList...)
		}
	}

	if 4&dataScope == 4 { //本部门数据权限
		structIds = append(structIds, ds.LoginData.StructId)
	}

	if 8&dataScope == 8 { //自定义
		for _, role := range ds.LoginData.RoleList {
			roleStructs := system.NewRoleStructService().GetRoleStructIdList(role)
			if len(roleStructs) > 0 {
				structIds = append(structIds, roleStructs...)
			}
		}
	}
	if 16&dataScope == 16 { //个人数据
		res.IsUser = true
	}

	structIds = tool.UniqueArrStr(structIds)
	res.StructIds = structIds

	return res
}

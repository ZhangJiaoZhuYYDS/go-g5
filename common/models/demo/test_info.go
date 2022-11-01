// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package demo

import (
	"b5gocmf/utils/core"
	"sync"
	"time"
)

type TestInfoModel struct {
	Id    string     `db:"id" json:"id" form:"id"` //  
	StructId    string     `db:"struct_id" json:"struct_id" form:"struct_id"` // 组织ID 
	UserId    string     `db:"user_id" json:"user_id" form:"user_id"` // 用户ID 
	Name    string     `db:"name" json:"name" form:"name"` // 名称 
	Status    string     `db:"status" json:"status" form:"status"` // 状态 
	Remark    string     `db:"remark" json:"remark" form:"remark"` // 介绍 
	CreateTime    *time.Time     `db:"create_time" json:"create_time" form:"-"` //  
	UpdateTime    *time.Time     `db:"update_time" json:"update_time" form:"-"` //  

}

func (m *TestInfoModel) Table() string {
	return "test_info"
}

// INew 给IModel接口使用创建一个新的结构体
func (m *TestInfoModel) INew() core.IModel {
	return m.New()
}

func (m *TestInfoModel) GetId() string {
	return m.Id
}

func (m *TestInfoModel) DataBase() string {
	return ""
}

func (m *TestInfoModel) New() *TestInfoModel {
	return &TestInfoModel{}
}

func (m *TestInfoModel) NewSlice() *[]TestInfoModel {
	return &[]TestInfoModel{}
}

var (
	instanceTestInfoModel *TestInfoModel //单例模式
	onceTestInfoModel     sync.Once
)

// NewTestInfoModel 单例获取
func NewTestInfoModel() *TestInfoModel {
	onceTestInfoModel.Do(func() {
		instanceTestInfoModel = &TestInfoModel{}
	})
	return instanceTestInfoModel
}

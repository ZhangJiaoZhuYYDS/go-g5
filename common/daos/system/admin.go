// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "b5gocmf/common/models/system"
	"sync"
)

type AdminDao struct {
	Model *AdminModel
}

var (
	instanceAdminDao *AdminDao //单例的对象
	onceAdminDao     sync.Once
)

func NewAdminDao() *AdminDao {
	onceAdminDao.Do(func() {
		instanceAdminDao = &AdminDao{Model: NewAdminModel()}
	})
	return instanceAdminDao
}

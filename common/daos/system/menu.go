// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package system

import (
	. "b5gocmf/common/models/system"
	"b5gocmf/utils/core"
	"sync"
)

type MenuDao struct {
	Model *MenuModel
}

var (
	instanceMenuDao *MenuDao //单例的对象
	onceMenuDao     sync.Once
)

func NewMenuDao() *MenuDao {
	onceMenuDao.Do(func() {
		instanceMenuDao = &MenuDao{Model: NewMenuModel()}
	})
	return instanceMenuDao
}

func (d *MenuDao) MenuTreeList() *[]MenuModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetField("id,parent_id,name").SetOrderBy(map[string]string{"parent_id": "asc", "list_sort": "asc", "id": "asc"}).Lists(list, "")
	return list
}
func (d *MenuDao) MenuLists()  *[]MenuModel {
	list := d.Model.NewSlice()
	_ = core.NewDao(d.Model).SetOrderBy(map[string]string{"parent_id": "asc", "list_sort": "asc", "id": "asc"}).Lists(list, "")
	return list
}
func (d *MenuDao) GetMenuShowLists(idList []string) *[]MenuModel {
	list := d.Model.NewSlice()
	where := "`type` != ?"
	args := make([]any, 0)
	args = append(args, "F")
	if idList != nil {
		where += " AND id in (?)"
		args = append(args, idList)
	}
	_= core.NewDao(d.Model).SetOrderBy(map[string]string{"parent_id": "asc", "list_sort": "asc", "id": "asc"}).Lists(list, where, args...)
	return list
}

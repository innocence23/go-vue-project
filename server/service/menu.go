package service

import (
	"errors"
	"project/entity"
	"project/zvar"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type MenuService struct {
}

func (menuService *MenuService) Create(menu entity.Menu) error {
	if !errors.Is(zvar.DB.Where("name = ?", menu.Name).First(&entity.Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return zvar.DB.Create(&menu).Error
}

func (menuService *MenuService) Update(menu entity.Menu) (err error) {
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	err = zvar.DB.Model(&entity.Menu{}).Where("id = ?", menu.ID).Updates(upDateMap).Error
	return err
}

func (menuService *MenuService) Find(id int) (err error, menu entity.Menu) {
	err = zvar.DB.First(&menu, id).Error
	return
}

func (menuService *MenuService) Hidden(id int) (err error) {
	err = zvar.DB.Model(&entity.Menu{}).Where("id = ?", id).Update("hidden", true).Error
	return
}

func (menuService *MenuService) Display(id int) (err error) {
	err = zvar.DB.Model(&entity.Menu{}).Where("id = ?", id).Update("hidden", false).Error
	return
}

//------ TreeList （列表）
func (menuService *MenuService) TreeList() (err error, list interface{}, total int64) {
	var menuList []entity.Menu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

func (menuService *MenuService) getBaseChildrenList(menu *entity.Menu, treeMap map[string][]entity.Menu) (err error) {
	menu.Children = treeMap[cast.ToString(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) getMenuTreeMap(roleId string) (err error, treeMap map[string][]entity.Menu) {
	var allMenus []entity.Menu
	treeMap = make(map[string][]entity.Menu)
	err = zvar.DB.Where("role_id = ?", roleId).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//------ end

//------ GetMenuTree （用户菜单）
func (menuService *MenuService) GetMenuTree(roleId string) (err error, menus []entity.Menu) {
	err, menuTree := menuService.getMenuTreeMap(roleId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]entity.Menu) {
	var allMenus []entity.Menu
	treeMap = make(map[string][]entity.Menu)
	err = zvar.DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func (menuService *MenuService) getChildrenList(menu *entity.Menu, treeMap map[string][]entity.Menu) (err error) {
	menu.Children = treeMap[cast.ToString(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//------ end

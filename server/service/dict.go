package service

import (
	"errors"
	"project/dto/request"
	"project/entity"
	"project/zvar"

	"gorm.io/gorm"
)

type DictService struct {
}

func (dictService *DictService) CreateDict(sysDictionary entity.Dict) (err error) {
	if (!errors.Is(zvar.DB.First(&entity.Dict{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = zvar.DB.Create(&sysDictionary).Error
	return err
}

func (dictService *DictService) DeleteDict(sysDictionary entity.Dict) (err error) {
	err = zvar.DB.Delete(&sysDictionary).Delete(&sysDictionary.DictDetails).Error
	return err
}

func (dictService *DictService) UpdateDict(sysDictionary *entity.Dict) (err error) {
	var dict entity.Dict
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := zvar.DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type == sysDictionary.Type {
		err = db.Updates(sysDictionaryMap).Error
	} else {
		if (!errors.Is(zvar.DB.First(&entity.Dict{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的type，不允许创建")
		}
		err = db.Updates(sysDictionaryMap).Error

	}
	return err
}

func (dictService *DictService) GetDict(Type string, Id int) (err error, sysDictionary entity.Dict) {
	err = zvar.DB.Where("type = ? OR id = ?", Type, Id).Preload("DictDetails").First(&sysDictionary).Error
	return
}

func (dictService *DictService) GetDictInfoList(info request.DictSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := zvar.DB.Model(&entity.Dict{})
	var sysDictionarys []entity.Dict
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return err, sysDictionarys, total
}

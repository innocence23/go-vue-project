package service

import (
	"project/dto/request"
	"project/entity"
	"project/zvar"
)

type DictDetailService struct {
}

func (dictDetailService *DictDetailService) CreateDictDetail(dictionaryDetail entity.DictDetail) (err error) {
	err = zvar.DB.Create(&dictionaryDetail).Error
	return err
}

func (dictDetailService *DictDetailService) DeleteDictDetail(dictionaryDetail entity.DictDetail) (err error) {
	err = zvar.DB.Delete(&dictionaryDetail).Error
	return err
}

func (dictDetailService *DictDetailService) UpdateDictDetail(dictionaryDetail *entity.DictDetail) (err error) {
	err = zvar.DB.Save(dictionaryDetail).Error
	return err
}

func (dictDetailService *DictDetailService) GetDictDetail(id int) (err error, dictionaryDetail entity.DictDetail) {
	err = zvar.DB.Where("id = ?", id).First(&dictionaryDetail).Error
	return
}

func (dictDetailService *DictDetailService) GetDictDetailList(info request.DictDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := zvar.DB.Model(&entity.DictDetail{})
	var dictionaryDetails []entity.DictDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.DictID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.DictID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dictionaryDetails).Error
	return err, dictionaryDetails, total
}

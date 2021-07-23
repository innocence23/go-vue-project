package service

import (
	"project/dto/request"
	"project/model/system"
	"project/zvar"
)

type OperationLogService struct {
}

func (operationRecordService *OperationLogService) Create(opLog system.OperationLog) (err error) {
	err = zvar.DB.Create(&opLog).Error
	return err
}

func (operationRecordService *OperationLogService) DeleteByIds(ids request.IdsReq) (err error) {
	err = zvar.DB.Delete(&[]system.OperationLog{}, "id in (?)", ids.Ids).Error
	return err
}

func (operationRecordService *OperationLogService) Delete(opLog system.OperationLog) (err error) {
	err = zvar.DB.Delete(&opLog).Error
	return err
}

func (operationRecordService *OperationLogService) Show(id uint) (err error, opLog system.OperationLog) {
	err = zvar.DB.Where("id = ?", id).First(&opLog).Error
	return
}

func (operationRecordService *OperationLogService) List(info request.OperationLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := zvar.DB.Model(&system.OperationLog{})
	var opLogs []system.OperationLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&opLogs).Error
	return err, opLogs, total
}

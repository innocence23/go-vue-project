package service

import (
	"project/dto/request"
	"project/entity"
	"project/zvar"
)

type OperationLogService struct {
}

func (operationRecordService *OperationLogService) Create(opLog entity.OperationLog) (err error) {
	err = zvar.DB.Create(&opLog).Error
	return err
}

func (operationRecordService *OperationLogService) DeleteByIds(ids []int) (err error) {
	err = zvar.DB.Delete(&[]entity.OperationLog{}, ids).Error
	return err
}

func (operationRecordService *OperationLogService) Delete(opLog entity.OperationLog) (err error) {
	err = zvar.DB.Delete(&opLog).Error
	return err
}

func (operationRecordService *OperationLogService) Show(id uint) (opLog entity.OperationLog, err error) {
	err = zvar.DB.Where("id = ?", id).First(&opLog).Error
	return
}

func (operationRecordService *OperationLogService) List(info request.OperationLogSearch) (opLogs []entity.OperationLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := zvar.DB.Model(&entity.OperationLog{})
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
	return
}

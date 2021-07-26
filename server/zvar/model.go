package zvar

import (
	"time"

	"gorm.io/gorm"
)

//todo 更改时间类型

type Model struct {
	ID        int            `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

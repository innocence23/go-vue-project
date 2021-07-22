package service

import (
	"project/service/autocode"
	"project/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

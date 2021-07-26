package service

import (
	"project/utils"
	"project/zvar"

	"go.uber.org/zap"
)

type MachineService struct {
}

func (ms *MachineService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		zvar.Log.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		zvar.Log.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		zvar.Log.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	return &s, nil
}

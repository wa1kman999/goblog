package system

import (
	"github.com/wa1kman999/goblog/pkg/system/model"
	"github.com/wa1kman999/goblog/pkg/system/service"
)

type ServiceSystem struct{}

func NewSystemService() *ServiceSystem {
	return &ServiceSystem{}
}

// GetServerInfo 服务器信息
func (app *ServiceSystem) GetServerInfo() (*model.Server, error) {
	cpu, err := service.InitCPU()
	if err != nil {
		return nil, err
	}
	ram, err := service.InitRAM()
	if err != nil {
		return nil, err
	}
	disk, err := service.InitDisk()
	if err != nil {
		return nil, err
	}
	return &model.Server{
		Os:   service.InitOS(),
		Cpu:  cpu,
		Ram:  ram,
		Disk: disk,
	}, nil
}

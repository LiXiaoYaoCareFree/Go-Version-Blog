package data_api

import (
	"Blog-Server/common/res"
	"Blog-Server/utils/computer"
	"github.com/gin-gonic/gin"
)

type ComputerDataResponse struct {
	CpuPercent  float64 `json:"cpuPercent"`
	MemPercent  float64 `json:"memPercent"`
	DiskPercent float64 `json:"diskPercent"`
}

func (DataApi) ComputerDataView(c *gin.Context) {
	var data = ComputerDataResponse{
		CpuPercent:  computer.GetCpuPercent(),
		MemPercent:  computer.GetMemPercent(),
		DiskPercent: computer.GetDiskPercent(),
	}
	res.OkWithData(data, c)
}

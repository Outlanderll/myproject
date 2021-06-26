package main

import (
	"awesomeProject1/pkg/setting"
	"awesomeProject1/routers"
	"awesomeProject1/underly"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init()  {
	setting.Setup()
	underly.Setup()
}
func main()  {
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf(":%d",setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.Run(endPoint)
}
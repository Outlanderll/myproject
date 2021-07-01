package main

import (
	"carrepair/dao"
	"carrepair/pkg/setting"
	"carrepair/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init()  {
	setting.Setup()
	dao.Setup()

}

func main()  {
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf(":%d",setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.Run(endPoint)
}

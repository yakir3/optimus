package main

import (
	"github.com/gin-gonic/gin"

	// "optimus-be/models"
	"optimus-be/routers"
	// "optimus-be/pkg/logging"
	// "optimus-be/pkg/setting"
	// "optimus-be/pkg/util"
)

// func init() {
// 	models.Setup()
// 	logging.Setup()
// 	util.Setup()
// }

// @title Optimus
// @version 0.1
// @description An example of gin
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)

	// routersInit := routers.InitRouter()
	// readTimeout := setting.ServerSetting.ReadTimeout
	// writeTimeout := setting.ServerSetting.WriteTimeout
	// endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	// endPoint := "127.0.0.1:8080"
	// maxHeaderBytes := 1 << 20

	// server := &http.Server{
	// 	Addr:    endPoint,
	// 	Handler: routersInit,
	// 	// ReadTimeout:    readTimeout,
	// 	// WriteTimeout:   writeTimeout,
	// 	MaxHeaderBytes: maxHeaderBytes,
	// }

	// log.Printf("[info] start http server listening %s", endPoint)

	// server.ListenAndServe()

	r := gin.Default()

	// load html template
	//r.LoadHTMLGlob("templates/*")

	// load static directory
	//r.Static("/static", "./static")

	//route group
	routers.InitRouter(r)

	r.Run()
}

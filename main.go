package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Es-war/go-gin-example/models"
	"github.com/Es-war/go-gin-example/pkg/logging"
	"github.com/Es-war/go-gin-example/pkg/setting"
	"github.com/Es-war/go-gin-example/routers"
	"github.com/gin-gonic/gin"
)

func main() {
    setting.Setup()
    models.Setup()
	logging.Setup()

    gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
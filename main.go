package main

import (
	"bcos_lottery_demo/app"
	"bcos_lottery_demo/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	// init route
	app.InitRouter(r)

	http.ListenAndServe(":"+strconv.Itoa(configs.Conf.Port), r)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Finished")
}

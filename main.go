/*
 * @Author: liumiu2020 liumiu2020@gmail.com
 * @Date: 2025-11-21 21:34:25
 * @LastEditors: liumiu2020 liumiu2020@gmail.com
 * @LastEditTime: 2025-11-22 12:05:09
 * @FilePath: \eth-client\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"flag"
	"fmt"
	"go-client/common"
	"go-client/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var env = flag.String("env", "", "app environment")

func main() {
	flag.Parse()
	err := godotenv.Load(".env" + *env)
	common.Init()
	//release resource
	defer release()
	fmt.Println(*env)
	if err != nil {
		fmt.Println("env error" + err.Error())
		return
	}
	route.R = gin.Default()
	route.RouteReg()
	var port string = "8888"
	if os.Getenv("port") != "" {
		port = os.Getenv("port")
	}
	route.R.Run(":" + port)
}

func release() {
	common.CommonClient.Close()
}

package main

import (
	"IM/router"
	"IM/utils"
)

// @title gin-swagger
// @version 1.0
// @description gin+gorm 测试swagger自动生成API文档
// @license.name Apache 2.0
// @host localhost:8080

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "127.0.0.1:8080")
}

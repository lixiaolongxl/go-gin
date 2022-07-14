package main

import (
	"order/dao"
	"order/models"
	"order/routers"
)

//开始拆分项目

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	DB := dao.InitMySQL()

	// 模型绑定
	DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	r.Run()
}

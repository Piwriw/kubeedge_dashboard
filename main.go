package main

import (
	"fmt"
	"new-ec-dashboard/controller"
	"new-ec-dashboard/dao/mysql"
	"new-ec-dashboard/logger"
	"new-ec-dashboard/router"
	"new-ec-dashboard/service"
	"new-ec-dashboard/setting"
	"os"
)

// @title 边缘计算管理平台项目接口文档
// @version 1.0
// @description Go web边缘计算管理平台项目接口文档


// @contact.name Piwriw
// @contact.url https://github.com/Piwriw
// @contact.email Piwriw@163.com

// @host 127.0.0.1:8000
func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: new-ec-dashboard config.yaml")
		return
	}
	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := service.InitK8sClient();err!=nil{
		fmt.Printf("init InitK8sClient failed, err:%v\n", err)
		return
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	//if err := redis.Init(setting.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}
	//defer redis.Close()


	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

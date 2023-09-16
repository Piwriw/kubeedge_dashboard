package router

import (
	"github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"new-ec-dashboard/logger"
	"new-ec-dashboard/middlewares"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "new-ec-dashboard/docs"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	//r.Use(logger.GinLogger(), logger.GinRecovery(true), m iddlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middlewares.Cors())
	//r.LoadHTMLFiles("./templates/index.html")
	//r.Static("/static", "./static")


	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//注册路由
	UserRouter(r)
	NodeRouter(r)
	PodRouter(r)
	DeviceRouter(r)
	RuleRouter(r)
	AIRouter(r)
	SvcDataRouter(r)
	IeRouter(r)
	ConfigMapRouter(r)
	NSRouter(r)
	SecretRouter(r)

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

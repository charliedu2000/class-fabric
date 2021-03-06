package internal

import (
	"github.com/gin-gonic/gin"
	"medicineApp/internal/config"
	"medicineApp/internal/middleware"
	"medicineApp/internal/router"
)

// InitGinEngine create gin engine
// Register router and middlewares
func InitGinEngine(r *router.Router) *gin.Engine {
	// 设置运行模式
	if config.C.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	// trace ID
	app.Use(middleware.TraceMiddleware())

	// copy body 因为 gin 的 reqbody 读取一次即不能使用所以进行复制
	app.Use(middleware.CopyBodyMiddleware())

	// 打印请求相关的 logger
	app.Use(middleware.LoggerMiddleware())

	// 恢复程序的中间件
	app.Use(middleware.RecoveryMiddleware())

	// 是否启动跨域访问
	if config.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	r.Register(app)

	return app
}

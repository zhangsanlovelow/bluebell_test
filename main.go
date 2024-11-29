package main

import (
	"bullbell_test/controller"
	"bullbell_test/dao/mysql"
	"bullbell_test/dao/redis"
	"bullbell_test/logger"
	"bullbell_test/pkg/snowflake"
	"bullbell_test/routes"
	"bullbell_test/settings"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1.初始化配置
	if err := settings.Init(); err != nil {
		zap.L().Error("init settings failed, err:", zap.Error(err))
		return
	}
	zap.L().Info("config init success...")
	//2.初始化日志
	if err := logger.Init(viper.GetString("mode")); err != nil {
		zap.L().Error("init logger failed, err:", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Info("logger init success...")

	//3.初始化mysql
	if err := mysql.Init(); err != nil {
		zap.L().Error("init mysql failed, err:", zap.Error(err))
		return
	}
	zap.L().Info("mysql init success...")
	defer mysql.Close()

	//4.初始化redis
	if err := redis.Init(); err != nil {
		zap.L().Error("init redis failed, err:", zap.Error(err))
		return
	}
	defer redis.Close()

	zap.L().Info("redis init success...")

	//4.初始化其他组件
	//...
	//...

	//初始化snowflake
	if err := snowflake.Init(viper.GetString("start_time"), viper.GetInt64("machine_id")); err != nil {
		zap.L().Error("init snowflake failed, err:", zap.Error(err))
		return
	}
	// fmt.Println(snowflake.GenID())

	//初始化validator
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("init validator failed, err:", zap.Error(err))
		return
	}

	//5.初始化路由
	r := routes.Setup()

	//6.启动服务
	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}

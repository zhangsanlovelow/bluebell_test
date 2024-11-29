package mysql

import (
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	// dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	dsn := viper.GetString("mysql.user") +
		":" + viper.GetString("mysql.password") +
		"@tcp(" + viper.GetString("mysql.host") +
		":" + viper.GetString("mysql.port") + ")/" +
		viper.GetString("mysql.db_name") +
		"?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))

		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

func Close() {
	_ = db.Close()
}

package mysql

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DefaultDB *gorm.DB
)

func Init() {
	dsn := viper.GetString("mysql.default.dsn")
	DefaultDB, _ = gorm.Open(mysql.Open(dsn))
}

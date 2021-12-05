package mysql

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

var (
	DefaultDB *gorm.DB
)

func Init() {
	dsn := viper.GetString("mysql.default.dsn")
	DefaultDB, _ = gorm.Open(mysql.Open(dsn))
}

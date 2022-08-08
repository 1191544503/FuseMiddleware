package initialize

import (
	"FuseMiddleware/global"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	// mysqlConfig := mysql.Config{
	// 	DSN:                       m.Dsn(), // DSN data source name
	// 	DefaultStringSize:         191,     // string 类型字段的默认长度
	// 	SkipInitializeWithVersion: false,   // 根据版本自动配置
	// }
	if db, err := gorm.Open("mysql", m.Dsn()); err != nil {
		return nil
	} else {
		db.SingularTable(true)
		db.LogMode(false)
		sqlDB := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

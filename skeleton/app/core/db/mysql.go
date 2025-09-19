package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewMySQL(dsn string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 赋值给全局变量
	Db = db

	return Db, nil
}

func GromGen(db *gorm.DB, path string) {
	// 创建query时，创建在指定数据库目录下
	g := gen.NewGenerator(gen.Config{
		OutPath:      path + "/query",
		ModelPkgPath: path + "/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})

	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		// g.GenerateModel("users"),
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

package main

import (
	"context"
	"gorm_app/models/model"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	db, err := NewDB()
	if err != nil {
		slog.Error("new db ", "error", err)
	}

	var resp []model.User
	err = db.WithContext(ctx).Where("id > ?", 100).Where("id < ?", 200).Find(&resp).Error
	if err != nil {
		slog.Error("find data", "error", err)
	}
	slog.Info("gorm select", "data", resp)

	// GromGen(db)
}

func NewDB() (*gorm.DB, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	return db, nil
}

func GromGen(db *gorm.DB) {
	// 创建query时，创建在指定数据库目录下
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./models/query",
		ModelPkgPath: "./models/model",
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

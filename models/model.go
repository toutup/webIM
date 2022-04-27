package models

import (
	"fmt"
	"log"
	"os"
	"time"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库实例
var db *gorm.DB

type Model struct {
	CreateAt time.Time `json:"createAt" db:"createAt"`
	UpdateAt time.Time `json:"updateAt" db:"updateAt"`
}

type Writer struct {
}

const (
	maxOpenConns = 320
	maxIdleConns = 32
	maxLifeTime  = 180
)

// 初始化数据库连接
func Init() {

	fmt.Println(">>>初始化mysql连接......")
	// 初始化mysql连接池
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	fmt.Println("dsn:", dsn)
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Microsecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	config := gorm.Config{
		SkipDefaultTransaction: true, // disable global write(create,update,delete) transaction, up 30% performance
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: newLogger,
	}

	db, err = gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		log.Fatalf("db models init err: %v", err)
	}

	mysqlDb, _ := db.DB()
	mysqlDb.SetMaxIdleConns(maxIdleConns)
	mysqlDb.SetMaxOpenConns(maxOpenConns)
	mysqlDb.SetConnMaxLifetime(maxLifeTime)

	fmt.Println(">>>mysql初始化成功......")
}

// 获取数据库连接
func GetDB() *gorm.DB {
	return db
}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

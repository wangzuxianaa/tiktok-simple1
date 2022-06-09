package model

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var RDB *redis.Client
var ctx = context.Background()

//
// MysqlInit
// @Description: Mysql数据库初始化
// @return err
//
func MysqlInit() (err error) {
	// 读取配置文件，配置文件为conf.yaml
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Conf.MysqlConfig.User,
		conf.Conf.MysqlConfig.Password,
		conf.Conf.MysqlConfig.Host,
		conf.Conf.MysqlConfig.Port,
		conf.Conf.MysqlConfig.Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return err
	}
	// 建立评论表，视频表，用户表和喜好表
	err = DB.AutoMigrate(&Comment{}, &Video{}, &User{}, &Favourite{})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//
// RedisInit
// @Description: Redis数据库初始化
// @return error
//
func RedisInit() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.RedisConfig.Addr,
		Password: conf.Conf.RedisConfig.Password,
		DB:       conf.Conf.RedisConfig.DB,
	})

	if _, err := RDB.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

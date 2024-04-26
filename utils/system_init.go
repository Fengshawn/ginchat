package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("config app inited...")
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("init MySQL error ...", err)
	} else {
		fmt.Println("Mysql inited...", DB)
	}

}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.PASSWORD"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})
	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("init Redis error ...", err)
	} else {
		fmt.Println("init Redis ...", pong)
		fmt.Println("Redis inited...")
	}

}

const (
	PublishKey = "websocket"
)

// Todo 之后改成kafka
// 发布消息到Redis
func Publish(ctx context.Context, channel string, payload string) error {
	//老的redis版本的Publish方法可能会用到 ctx, 高版本的只用channel和msg
	var err error
	fmt.Println("Publish>>>>>>>>>>", payload)
	err = Red.Publish(channel, payload).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	//老的redis版本的PSubscribe方法可能会用到 ctx, 高版本的只用channel和msg
	sub := Red.Subscribe(channel)
	fmt.Println("Subscribe ctx ...", ctx)
	msg, err := sub.ReceiveMessage()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe ...", msg.Payload)
	return msg.Payload, err
}

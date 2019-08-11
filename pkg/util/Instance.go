package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/viletyy/potato/pkg/logging"
	"github.com/viletyy/potato/pkg/setting"
	"time"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt	*time.Time	`sql:"index" json:"deleted_at"`
}


var(
	DB = InitDB()
	Redis *redis.Client
)


func init()  {
	// 数据库配置
	InitDB()
	//defer DB.Close()
	InitRedis()
	defer Redis.Close()
}

func InitDB() *gorm.DB {
	var (
		err error
		dbType, dbName, user, password, host, port, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Info(2, fmt.Sprintf("Fail to get section 'database': %v", err))
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	DB, err := gorm.Open(dbType, fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", host, user, dbName, port, password))

	if err != nil {
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB.LogMode(true)

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	logging.Info(DB)

	return DB
}

func InitRedis() {
	sec, err := setting.Cfg.GetSection("redis")

	if err != nil {
		logging.Info(2, fmt.Sprintf("Fail to get section 'redis': %v", err))
	}

	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			sec.Key("HOST").String(),
			sec.Key("PORT").String(),
		),
	})

	if _, err := Redis.Ping().Result(); err != nil {
		logging.Fatal("redis连接失败!", err)
	}
}

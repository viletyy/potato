package util

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/viletyy/potato/pkg/logging"
	"github.com/viletyy/potato/pkg/setting"
	"reflect"
	"strconv"
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

func FillStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值
	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}

func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

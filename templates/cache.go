package templates

import "fmt"

func RedisGo(moduleName string) string {
	return fmt.Sprintf(`package cache

import (
	"context"
	"time"

	"%s/conf"
	"%s/log"
	"github.com/redis/go-redis/v9"
)

var RedisCli *redis.Client

func NewCache() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.Addr,
		Password: conf.Conf.Redis.Pwd,
		DB:       conf.Conf.Redis.Db,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}

	RedisCli = rdb
	log.Zlog.Info("Redis connected")
	return nil
}

func SetValue(key string, data interface{}) error {
	_, err := RedisCli.Set(context.Background(), key, data, 0).Result()
	return err
}

func GetValue(key string) string {
	str, err := RedisCli.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		}
		log.Zlog.Errorf("GetValue err: %%+v", err)
		return ""
	}
	return str
}

func SetValueWithTTL(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	_, err := RedisCli.Set(ctx, key, data, ttl).Result()
	return err
}

func DeleteKey(key string) {
	RedisCli.Del(context.Background(), key)
}

func IncrValue(key string) error {
	_, err := RedisCli.Incr(context.Background(), key).Result()
	return err
}
`, moduleName, moduleName)
}

func KeysGo() string {
	return `package cache

// Redis key definitions
const (
	CacheUserToken = "user:token:%s"      // User token
	CacheUserInfo  = "user:info:%s"       // User info
)
`
}

func DbGo(moduleName string) string {
	return fmt.Sprintf(`package storage

import (
	"%s/conf"
	"%s/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.Zlog.Info("DB start initialize")

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.Conf.MysqlAddr,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		log.Zlog.Fatalf("gorm open err: %%+v", err)
	}

	log.Zlog.Info("DB initialize success")
	DB = gormDB
}
`, moduleName, moduleName)
}

func UserModelGo() string {
	return `package storage

import "gorm.io/gorm"

type User struct {
	ID         int64  ` + "`gorm:\"primaryKey;autoIncrement\"`" + `
	UID        string ` + "`gorm:\"column:uid;type:varchar(64);uniqueIndex\"`" + `
	Nickname   string ` + "`gorm:\"column:nickname;type:varchar(128)\"`" + `
	Avatar     string ` + "`gorm:\"column:avatar;type:varchar(512)\"`" + `
	Phone      string ` + "`gorm:\"column:phone;type:varchar(20);index\"`" + `
	Status     int    ` + "`gorm:\"column:status;default:1\"`" + ` // 1:active 0:disabled
	CreateTime int64  ` + "`gorm:\"column:create_time;autoCreateTime\"`" + `
	UpdateTime int64  ` + "`gorm:\"column:update_time;autoUpdateTime\"`" + `
}

func (User) TableName() string {
	return "t_user"
}

func GetUserByUID(uid string) (*User, error) {
	var user User
	err := DB.Where("uid = ?", uid).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByPhone(phone string) (*User, error) {
	var user User
	err := DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func UpdateUser(uid string, updates map[string]interface{}) error {
	return DB.Model(&User{}).Where("uid = ?", uid).Updates(updates).Error
}
`
}

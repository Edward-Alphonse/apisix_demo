package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(config *Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         config.Addr(),
		Username:     config.User,
		Password:     config.Password,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
		MinIdleConns: config.MinIdleConns,
	})
}

func GetRedis() *redis.Client {
	return rdb
}

//
//func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
//	err := rdb.Set(ctx, key, value, expiration).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func Get(ctx context.Context, key string) (string, error) {
//	val2, err := rdb.Get(ctx, key).Result()
//	return val2, err
//}
//
//func IsNilErr(err error) bool {
//	return err == redis.Nil
//}
//
//func GetJsonWithFunc(ctx context.Context, key string, ttl time.Duration, result interface{}, getter func() (interface{}, error)) error {
//	var err error
//	getResult := rdb.Get(ctx, key)
//	bytes, _ := getResult.Bytes()
//	if len(bytes) > 0 {
//		// 从缓存获得结果
//		err = json.Unmarshal(bytes, result)
//		if err == nil {
//			return nil
//		}
//		logs.ErrorByArgs("[redis] failed to unmarshal data:%s", string(bytes))
//	}
//	// 执行函数，获得结果
//	r, err := getter()
//
//	if err == nil && !isNil(r) {
//		bytes, err = json.Marshal(r)
//		if err != nil {
//			logs.ErrorByArgs("[redis] json.Marshal failed:%s", err)
//		}
//		if bytes != nil && string(bytes) != "null" {
//			// 写入缓存
//			rdb.SetEx(ctx, key, bytes, ttl)
//			_ = json.Unmarshal(bytes, result)
//		}
//	}
//	return err
//}
//func isNil(i interface{}) bool {
//	vi := reflect.ValueOf(i)
//	if vi.Kind() == reflect.Ptr {
//		return vi.IsNil()
//	}
//	return false
//}

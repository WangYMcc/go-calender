package redis

import (
	"core/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

var(
	pool *redis.Pool
	redisHost = beego.AppConfig.String("redis.host") + ":" + beego.AppConfig.String("redis.port")
	redisPass = beego.AppConfig.String("redis.password")
	isCacheInit bool
)

func newRedisPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle: 50,
		MaxActive: 30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn,  error) {
			c,err := redis.Dial("tcp",
				redisHost,
				redis.DialPassword(redisPass))
			if err != nil {
				beego.Error(err)
				return nil,err
			}
			//2、访问认证
			//if _, err =c.Do("AUTH",redisPass);err!=nil{
			//	c.Close()
			//	return nil,err
			//}
			return c,nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_,err := conn.Do("PING")
			return err
		},
	}
}

func init(){
	pool = newRedisPool()
	isCacheInit = true
	beego.Debug("init")
}

//对外暴露连接池
func GetRedisPool() *redis.Pool{
	return pool
}


func SET(key string, val interface{}) error{
	conn := pool.Get()
	//注意close()
	defer conn.Close()


	var u []byte
	var err error
	u, err = json.Marshal(val)

	if err != nil {
		return err
	}

	_, e := conn.Do("SET", key, u)

	return e
}

func GET(key string, val interface{}) (models.Model, error){
	conn := pool.Get()
	//注意close()
	defer conn.Close()


	var u []byte
	var err error
	u, err = json.Marshal(val)

	if err != nil {
		return nil, err
	}

	_, e := conn.Do("SET", key, u)

	return nil, e
}

func IsRedisCacheInit() bool{
	return isCacheInit
}

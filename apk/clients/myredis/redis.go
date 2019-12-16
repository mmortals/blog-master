package myredis

import (
	"errors"
	"github.com/go-redis/redis"
)

type ACC_TYPE string

const (
	MASTER                ACC_TYPE = "w"
	SLAVE                 ACC_TYPE = "r"
	DEFAULT_POOL_SIZE     int      = 10 //默认最大连接数
	DEFAULT_MIN_IDLE_CONN int      = 5  //默认闲置连接数
)

// 单例管理器
var MyfRedis *redis.Client

// 单实例配置
type CACHEConnConf struct {
	Host string `toml:"host"` // host+端口
	Port int    `toml:"port"` // 端口
}

// Redis主从配置
type CACHEGroupConf struct {
	Name     string `toml:"name"`
	Password string `toml:"password"` // 密码

	//Master *CACHESubGroupConf `toml:"Master"` // 主库配置
	//Slaves *CACHESubGroupConf `toml:"Slave"`  // 从库配置

}

//多redis实例，配置
type CACHEConf struct {
	GroupConfList []CACHEGroupConf `toml:"Group"`
}

// 错误码
var (
	ERR_CACHE_NAME_NOT_FOUND  = errors.New("redis数据库名称不能为空")
	ERR_CACHE_GROUP_NOT_FOUND = errors.New("此redis实例不存在")
	ERR_CACHE_CONN_NOT_FOUND  = errors.New("没有可用redis实例连接")
)

func ConnectRedis() (err error) {

	MyfRedis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = MyfRedis.Ping().Result()

	return
}

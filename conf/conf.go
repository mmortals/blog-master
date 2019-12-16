package conf

import (
	"blog-master/apk/clients/myredis"
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	confPath *string
	MyfConf  *MyfConfig
)

// 框架配置实例
type MyfConfig struct {
	Domain string             `toml:"domain"` // 域名
	Debug  int                `toml:"debug"`  // 调试模式
	Redis  *myredis.CACHEConf `toml:"Redis"`  //redis配置
}

// 初始化配置
func InitConf() (err error) {
	MyfConf = defaultConf()
	if _, err = toml.DecodeFile(*confPath, MyfConf); err != nil {
		return
	}
	return
}

// 默认配置
func defaultConf() *MyfConfig {
	return &MyfConfig{
		Debug:  0,
		Domain: "",
		Redis:  nil,
	}
}

func init() {
	confPath = flag.String("conf", "./app.toml", "框架配置文件路径")
}

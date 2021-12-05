package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wa1kman999/goblog/pkg/common/constants"
)

var globalConfig Config

// Config 全局配置
type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

// Init 初始化配置
func Init(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constants.ConfigEnv); configEnv == "" {
				config = constants.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", constants.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&globalConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&globalConfig); err != nil {
		fmt.Println(err)
	}
	return v
}

// Get 获取配置
func Get() Config {
	fmt.Printf("打印配置:%#v", globalConfig)
	return globalConfig
}

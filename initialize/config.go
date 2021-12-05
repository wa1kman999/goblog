package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wa1kman999/goblog/global"
)

// ConfigInit 初始化配置
func ConfigInit() error {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GBConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GBConfig); err != nil {
		return err
	}
	return nil
}

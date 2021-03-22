/*
 * @Date: 2021-03-22 09:38:24
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 10:12:33
 * @FilePath: /potato/initialize/viper.go
 */
package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/viletyy/potato/global"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GO_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GO_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}

package pkg

import (
	"fmt"
	"path/filepath"
	_ "project/packfile"
	"project/zvar"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(config string) *viper.Viper {
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
		if err := v.Unmarshal(&zvar.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&zvar.Config); err != nil {
		fmt.Println(err)
	}
	zvar.Config.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

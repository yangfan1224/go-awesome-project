package viper

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/fsnotify/fsnotify"
)

func init(){
	//viper.SetConfigFile("/home/yangfan/go_workspace/src/sndo.com/awesomeProject1/viper/conf.toml")
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func ReadConfig() string {
	author := viper.GetString("author")
	fmt.Print(author)
	return author
}
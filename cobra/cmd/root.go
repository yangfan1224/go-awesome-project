package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var author string

func init()  {
	cobra.OnInitialize(initConfig)
}

func NewRootCmd() *cobra.Command{
	rootCmd := &cobra.Command{
		Use: "testserver",
		Short: "testserver is a test command for cobra",
		Long: `testserver is a test command for cobra`,
		//Args: func(cmd *cobra.Command, args []string) error {
		//	err := cobra.MinimumNArgs(1)(cmd, args)
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//},
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {

			config, _:= cmd.PersistentFlags().GetString("config")
			fmt.Printf("config is %v \n", config)
		},
	}

	rootCmd.PersistentFlags().StringP( "author", "a", "", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringP("config","c","","config file path")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))


	return rootCmd
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	viper.SetConfigFile("/home/yangfan/go_workspace/src/sndo.com/awesomeProject1/cobra/conf.toml")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

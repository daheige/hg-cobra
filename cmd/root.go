package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//根命令设置
var rootCmd = &cobra.Command{
	Use: "hg-cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("name: ", viper.GetString("name"))

		fmt.Println("OK")
		fmt.Println("author: ", viper.GetString("author"))
		fmt.Println("userInfo", viper.Get("userInfo").(*User))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var name string
var age int

type User struct {
	Id   int
	Name string
}

var cfgFile string

func init() {
	//用于viper读取配置，而viper之所以能直接用，是因为底层源码init()方法实例化了var v *Viper
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "Person's age")

	//将rootCmd接收的参数绑定到viper
	rootCmd.Flags().StringP("author", "u", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.Flags().Lookup("author"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("userInfo", &User{
		Id:   1,
		Name: "daheige",
	})

}

// initConfig reads in config file and ENV variables if set.
// 初始化配置文件路径
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		//     fmt.Println(err)
		//     os.Exit(1)
		// }

		// Search config in home directory with name ".env" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName("app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

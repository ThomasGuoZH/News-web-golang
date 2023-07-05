package main

import (
	"backend/mysql"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//"github.com/sirupsen/logrus"

	"backend/config"
	"backend/models"
	_ "backend/mysql"
	"backend/router"
)

var (
	cfgFile string
	//logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func main() {
	if err := Execute(); err != nil {
		log.Println("start fail:", err.Error())
		os.Exit(-1)
	}
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := mysql.Init(
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		mysql.DB.AutoMigrate(&models.User{})
		mysql.DB.AutoMigrate(&models.News{})
		mysql.DB.AutoMigrate(&models.Comment{})
		mysql.DB.AutoMigrate(&models.Likes{})
		mysql.DB.AutoMigrate(&models.UserFav{})

		defer mysql.DB.Close()

		r := router.SetupRouter()
		err = r.Run()
		if err != nil {
			return err
		}

		port := viper.GetString("port")
		log.Println("port = *** =", port)
		return http.ListenAndServe(port, nil) // listen and serve
	}

	return rootCmd.Execute()
}

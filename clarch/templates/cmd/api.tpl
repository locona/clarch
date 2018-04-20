package cmd

import (
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/infra"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project/logger"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project/middleware"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project/validator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ListenApi()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	project.InitConfig()
	logger.InitLogger()
	validator.InitValidator()
	infra.InitMysql()
	infra.InitRedis()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListenApi() {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// middleware
	{
		middle := middleware.NewMiddleware()
		api.Use(middle.Cors())
		api.Use(middle.Logging())
	}

	infra.Mysql.LogMode(true)
	api.Run(":" + project.Config.API.Port)
}

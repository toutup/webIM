package apiServer

import (
	"fmt"
	"xiliangzi_pro/internal/pkg/conf"
	"xiliangzi_pro/models"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "apiServer",
	Short: "start api server",
	Long:  "start api server",
	Run: func(cmd *cobra.Command, args []string) {
		handle()
	},
}

var (
	port int
	env  string
)

// 获取环境配置
func init() {
	Cmd.Flags().IntVar(&port, "port", 8000, "api server port")
	Cmd.Flags().StringVar(&env, "env", "local", "envionment variables")
}

func handle() {
	fmt.Println("port:", port)
	// 加载配置文件
	conf.Init()
	// 初始化数据库连接
	models.Init()
	// 初始化api server
	ApiServerInit(port)
}

package websocket

import (
	"xiliangzi_pro/internal/pkg/conf"
	"xiliangzi_pro/models"
	"xiliangzi_pro/pkg/redisConn"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "websocket",
	Short: "start websocket server",
	Long:  "start websocket server",
	Run: func(cmd *cobra.Command, args []string) {
		handle()
	},
}

var (
	addr string
	port int
	env  string
)

// 获取命令配置
func init() {
	Cmd.Flags().StringVar(&addr, "addr", "127.0.0.1", "websocket addr")
	Cmd.Flags().IntVar(&port, "port", 9001, "websocket port")
	Cmd.Flags().StringVar(&env, "env", "dev", "envionment variables")
}

func handle() {
	// 加载配置
	conf.Init()
	// 初始化数据库
	models.Init()
	// 初始化redis
	redisConn.Init()
	// 初始化websocket server
	WsServerInit(port)
}

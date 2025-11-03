package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/buildingBlocksCoder/writeCode"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "Coder", Short: "生成代码CLI", Version: "1.0"}
	cmds := []*cobra.Command{
		{
			Use:     "start",
			Short:   "start the corder",
			Example: "./coder start backend || ./coder start front",
			Run: func(cmd *cobra.Command, args []string) {

				// 第一个参数必填
				if len(args) < 1 {
					fmt.Println("请输入[backend,front]参数区分后端还是前端")
					return
				}

				// 后端端口 2402 前端端口 2403
				type_ := args[0]
				port := "2402"
				if type_ == "backend" {
					fmt.Println("后端代码助手启动中...")
				} else {
					port = "2403"
					fmt.Println("前端代码助手启动中...")
				}

				gin.SetMode(gin.ReleaseMode)

				// 实例化引擎
				r := gin.Default()

				// 支持跨域
				r.Use(writeCode.Cors())

				// 健康检查
				r.GET("/health", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "pong",
					})
				})

				// 生成代码
				r.POST("/writeCode", writeCode.WriteCode)

				// 启动服务
				fmt.Println("欢迎使用Coder")
				r.Run(fmt.Sprintf("localhost:%s", port))
			},
		},
	}
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

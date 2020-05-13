package main

import (
	"github.com/huangxinchun/hxcgo/admin"
	"github.com/spf13/cobra"
)

//test
func main() {
	var modelList = &cobra.Command{Use: "god"}

	//后台服务
	modelList.AddCommand(admin.AdminModel)

	//执行命令
	modelList.Execute()
}

package cmd

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/cobra"
	"jettjia/igo-ddd-cli/cmd/build"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build go project",
	Long:  `build go project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			helpBuild()
			return
		}
		//处理传入的参数
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			helpBuild()
			return
		}
		build.Run(args, name)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// 接受传入的参数：name， 短拼是：n， 默认值是：main， 提示信息是： 编译后文件名字
	buildCmd.Flags().StringP("name", "n", "main", "output file name")
}

func helpBuild() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    go-ddd-cli build [OPTION]

ARGUMENT
    OPTION
	-n	output file name

EXAMPLES
    go-ddd-cli build main.go -n main
`))
}

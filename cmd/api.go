package cmd

import (
	"cobra/interfaces"

	"github.com/spf13/cobra"
)

var quickStart string
var port string

func init() {
	rootCmd.AddCommand(api)
	api.Flags().StringVarP(&quickStart, "start", "s", "", "quick start")
	//api.Flags().StringVarP(&quickStart, "start", "s", "", "quick start")//在一个命令下相同的flag会panic
	api.Flags().StringVarP(&port, "port", "p", ":8080", "set port")
}

var api = &cobra.Command{
	Use:   "api",
	Short: "启动web服务",
	Run: func(cmd *cobra.Command, args []string) {
		if port != ":8080" {
			if port[0] != ':' {
				port = ":" + port
			}
		}

		interfaces.Router(port)
	},
}

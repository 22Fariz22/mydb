package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mydb/cli/mydb/pkg"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get value via key",
	Long: `command get provide one argument: key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		value,err:= pkg.Get(args)
		if err!=nil{
			log.Fatal("error in get.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mydb/cli/mydb/pkg"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a pair",
	Long: `delete a pair via key. command get provide one arguments: key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("del called")

		_,err:= pkg.Del(args)
		if err!=nil{
			log.Fatal("error in del.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("Успешно удалено")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mydb/cli/mydb/pkg"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set key and value",
	Long: `provide two arguments: key and value`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")

		_,err:= pkg.Set(args)
		if err!=nil{
			log.Fatal("error in set.go: ",err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("Успешно сохранено")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

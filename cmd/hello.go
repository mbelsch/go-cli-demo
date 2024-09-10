package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Hello struct {
	message string
}

var helloParams = Hello{}

func initHello(parent *cobra.Command) {
	parent.AddCommand(helloCmd)
	helloCmd.Flags().StringVarP(&helloParams.message, "message", "m", "You didnt provide a message. Try to execute 'helper hello -m Awesome CLI'", "The message, that should be printed")
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello all services to specified environment",
	Long:  `hello all services to specified environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(helloParams.message)
	},
}

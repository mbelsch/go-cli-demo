package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helper",
	Short: "This cli is platform independent cli",
	Long: `This cli is platform independent cli. 
			It uses Go and spf13/cobra to rapidly build new commands`,
	Run: nil,
}

func initRoot() {
	initHello(rootCmd)
	initLogin(rootCmd)
}

func Execute() {
	initRoot()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

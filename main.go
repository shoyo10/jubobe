package main

import (
	"fmt"
	"jubobe/cmd"
	"os"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "jubobe"}

func main() {
	rootCmd.AddCommand(cmd.ServerCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute root command failed, err:%+v", err)
		os.Exit(1)
	}
}

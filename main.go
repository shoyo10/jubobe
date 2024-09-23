package main

import (
	"fmt"
	"jubobe/cmd/apiserver"
	"jubobe/cmd/pgmigration"
	"os"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "jubobe",
}

func main() {
	rootCmd.AddCommand(apiserver.ServerCmd, pgmigration.ServerCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute root command failed, err:%+v", err)
		os.Exit(1)
	}
}

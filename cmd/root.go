/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "line-notify",
	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "gofiber-creator",
	Short:   "Go Fiber project scaffolding tool",
	Long:    `Generate Go project template based on Fiber + GORM + Redis`,
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("gofiber-creator version {{.Version}}\n")
}

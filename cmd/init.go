package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/halfhuman88/gofiber-creator/templates"
	"github.com/spf13/cobra"
)

var (
	projectName string
	moduleName  string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go Fiber project",
	Long:  `Create a complete Go Fiber project structure in current directory`,
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Project name (required)")
	initCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Go module name (e.g. github.com/user/project)")
	initCmd.MarkFlagRequired("name")
}

func runInit(cmd *cobra.Command, args []string) {
	if projectName == "" {
		fmt.Println("Error: please specify project name with --name")
		os.Exit(1)
	}

	if moduleName == "" {
		moduleName = projectName
	}

	fmt.Printf("🚀 Creating project: %s\n", projectName)
	fmt.Printf("📦 Module: %s\n", moduleName)

	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		os.Exit(1)
	}

	// Generate all files
	files := templates.GetAllTemplates(moduleName, projectName)
	for path, content := range files {
		fullPath := filepath.Join(projectName, path)
		dir := filepath.Dir(fullPath)

		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			continue
		}

		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			fmt.Printf("Failed to write file %s: %v\n", fullPath, err)
			continue
		}
		// Simplify path display
		displayPath := strings.TrimPrefix(fullPath, projectName+"/")
		fmt.Printf("  ✅ %s\n", displayPath)
	}

	fmt.Println("\n🎉 Project created successfully!")
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  # Edit dev.yaml config\n")
	fmt.Printf("  go run cmd/main.go -config_path=dev.yaml\n")
}

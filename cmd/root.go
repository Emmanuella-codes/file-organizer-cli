package cmd

import (
	"file-organizer-cli/internal/config"
	"file-organizer-cli/internal/logger"
	"file-organizer-cli/internal/organizer"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	logDir  string
)

var rootCmd = &cobra.Command{
	Use: "file-organizer",
	Short: "A CLI tool to organize files into folders",
	Long: `File Organizer is a CLI tool that helps you organize files in a directory 
		by grouping them into categories such as Documents, Images, Videos, etc..`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetDir := "."
		if len(args) > 0 {
			targetDir = args[0]
		}

		log, err := logger.New(logDir)
		if err != nil {
			return err
		}
		defer log.Close()

		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			return err
		}

		org := organizer.New(targetDir)

		if len(cfg.CustomCategories) > 0 {
			customCats := make([]organizer.FileCategory, 0)
			for name, extensions := range cfg.CustomCategories {
					customCats = append(customCats, organizer.FileCategory{
							Name:       name,
							Extensions: extensions,
					})
			}
			org.Categories = customCats
		}

		if err := org.Organize(); err != nil {
			log.Log("Error organizing files: " + err.Error())
			return err
		}

		log.Log("Successfully organized files in " + targetDir)
    return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
			homeDir = "."
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", 
			filepath.Join(homeDir, ".file-organizer.json"), "config file path")
	rootCmd.PersistentFlags().StringVar(&logDir, "log-dir", 
			filepath.Join(homeDir, ".file-organizer", "logs"), "log directory path")
}

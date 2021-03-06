package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var baseDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goatd",
	Short: "goatd is a GTD (Getting Things Done) task manager",
	Long: `goatd stores all your GTD data in plain-text MarkDown .md files and
allows you to work with it via the command line or via the web browser.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		"config file (default is /etc/goatd/goatd.yml or $HOME/.config/goatd/goatd.yaml)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&baseDir,
		"base-dir",
		"b",
		"",
		"base directory containing everything (default is $HOME/goatd)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("baseDir", rootCmd.PersistentFlags().Lookup("base-dir"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath("/etc/goatd/")
		viper.AddConfigPath(filepath.Join(home, ".config", "goatd"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("goatd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

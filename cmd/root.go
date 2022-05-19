package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goatd",
	Short: "goatd is a GTD (Getting Things Done) task manager",
    Long: `goatd stores all your GTD data in plain-text MarkDown .md files and
allows you to work with it via the command line or via the web browser.

To initialize a new GTD folder structure:

    goatd init  # Not implemented

To verify the structure and content of the MarkDown files (e.g. after manually
modifying them):

    goatd verify  # Not implemented

To run the webserver:

    goatd server  # Not implemented

To interact with projects:

    goatd projects  # Not implemented

To interact with tasks:

    goatd tasks  # Not implemented

To easily capture (i.e. add to inbox) a new task:

    goatd capture  # Not implemented
`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goatd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

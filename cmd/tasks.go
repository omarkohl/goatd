package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func initBaseDir() error {
	// This initialization for the baseDir should go somewhere else,
	// but it's not entirely clear to me where the best place is
	// Maybe PersistentPreRun .
	if baseDir == "" {
		var home string
		var err error
		if home, err = os.UserHomeDir(); err != nil {
			return fmt.Errorf("can't get home dir: %w", err)
		}
		baseDir = path.Join(home, "goatd")
	}
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("can't create baseDir %v: %w", baseDir, err)
	}
	dirs := []string{
		"projects",
		"someday",
		"trash",
		"done",
	}
	for _, d := range dirs {
		d = path.Join(baseDir, d)
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("can't create dir %v: %w", d, err)
		}
	}
	inboxPath := path.Join(baseDir, "inbox.md")
	f, err := os.OpenFile(inboxPath,
		os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("can't create %v: %w", inboxPath, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("can't close %v: %w", inboxPath, err)
	}
	return nil
}

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Display all open tasks sorted by priority",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initBaseDir(); err != nil {
			return err
		}
		fmt.Println(baseDir)
		return nil
	},
}

func init() {
	listCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

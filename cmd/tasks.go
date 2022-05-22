package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

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
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		fmt.Printf("Will create GTD structure under %v\n", baseDir)
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

func getTasks(baseDir string) ([]string, error) {
	var tasks []string

	fmt.Println("Printing whole content of project files instead of only" +
		" tasks. Continue working here...")

	err := filepath.Walk(
		path.Join(baseDir, "projects"),
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Println(path)
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return fmt.Errorf("could not read %v: %w", path, err)
				}
				fmt.Println(content)
			}
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %w", err)
	}

	return tasks, nil
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
		fmt.Printf("Working with baseDir: %v\n", baseDir)
		tasks, err := getTasks(baseDir)
		if err != nil {
			return err
		}
		for _, t := range tasks {
			fmt.Println(t)
		}
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

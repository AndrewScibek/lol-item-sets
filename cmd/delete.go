package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [Champions Directory]",
	Short: "Deletes old item sets",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		deleteSets(path)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteSets(dirPath string) {
	var input string
	fmt.Print("This will delete all files in ", dirPath, "\nAre you sure you want to do that? (Y/n): ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input != "" && !strings.HasPrefix(input, "y") && !strings.HasPrefix(input, "n") {
		fmt.Println("Invalid input")
		return
	}
	if strings.HasPrefix(input, "n") {
		return
	}
	fmt.Println("Deleting files")

	err := removeContents(dirPath)
	if err != nil {
		fmt.Println("Couldn't delete all sets")
	}
}
func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

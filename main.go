package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "searchex",
		Short: "Search exploits",
	}

	newCmd := &cobra.Command{
		Use:   "new",
		Short: "Show new",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Println("Additional argument:", args[0])
			} else {
				fmt.Println("No additional arguments")
			}
		},
	}

	rootCmd.AddCommand(newCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

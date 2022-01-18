package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// Version Info
	Version = "1.0.0"
)

var (
	StartVCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: AppName + " version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}
)

func runV() error {
	fmt.Println(Version)
	return nil
}

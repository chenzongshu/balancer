package app

import (
"fmt"

"github.com/spf13/cobra"
"balancer/pkg/version"
)

func NewVersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version of balancer",
		Long:  `Prints the version of balancer.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("balancer version %+v\n", version.Get())
		},
	}
	return versionCmd
}

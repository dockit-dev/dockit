/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dockit-dev/dockit/internal/command/configure"
	dcontext "github.com/dockit-dev/dockit/internal/command/docker/context"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure [path]",
	Short: "Configures access to a remote dockit instance",
	Long: `The 'configure' command sets up access to a remote docker server
by providing the dockit configuration file. It creates a new docker context
and sets it as active, enabling seamless interaction with the remote Dockit instance.`,
	Example:               "  dockit configure /path/to/gz.tar",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Configuring access to remote Dockit instance from file: %s\n", args[0])

		if err := configure.Run(args[0]); err != nil {
			return err
		}

		if err := dcontext.Create(); err != nil {
			return err
		}

		fmt.Println("\nDockit configuration is set up successfully!")
		fmt.Println("You can now use Docker CLI to interact with the remote Dockit instance.")
		fmt.Println("\nExample:")
		fmt.Println("  docker ps -a")
		fmt.Println("\nTo switch back to using local Docker, set the context to the default one:")
		fmt.Println("  docker context use default")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
